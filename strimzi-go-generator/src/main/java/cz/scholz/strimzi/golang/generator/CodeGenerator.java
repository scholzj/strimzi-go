package cz.scholz.strimzi.golang.generator;

import com.fasterxml.jackson.databind.ObjectMapper;
import io.fabric8.kubernetes.api.model.Quantity;
import io.strimzi.api.annotations.ApiVersion;
import io.strimzi.api.kafka.model.bridge.KafkaBridge;
import io.strimzi.api.kafka.model.common.template.PodDisruptionBudgetTemplate;
import io.strimzi.api.kafka.model.connect.KafkaConnect;
import io.strimzi.api.kafka.model.connector.KafkaConnector;
import io.strimzi.api.kafka.model.kafka.Kafka;
import io.strimzi.api.kafka.model.kafka.SingleVolumeStorage;
import io.strimzi.api.kafka.model.kafka.listener.GenericKafkaListener;
import io.strimzi.api.kafka.model.mirrormaker2.KafkaMirrorMaker2;
import io.strimzi.api.kafka.model.nodepool.KafkaNodePool;
import io.strimzi.api.kafka.model.rebalance.KafkaRebalance;
import io.strimzi.api.kafka.model.topic.KafkaTopic;
import io.strimzi.api.kafka.model.user.KafkaUser;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;

import java.io.FileNotFoundException;
import java.io.FileOutputStream;
import java.io.IOException;
import java.io.OutputStreamWriter;
import java.lang.reflect.InvocationTargetException;
import java.lang.reflect.Method;
import java.lang.reflect.ParameterizedType;
import java.nio.charset.StandardCharsets;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Locale;
import java.util.Map;
import java.util.Set;
import java.util.Stack;

public class CodeGenerator {
    private final static Logger LOGGER = LogManager.getLogger(CodeGenerator.class);

    private static final String NL = System.lineSeparator();
    private static final String TAB = "    ";

    private static final Path HEADER_BOILERPLATE_PATH = Path.of("../hack/boilerplate.go.txt");
    private static final String OUTPUT_PATH = "../pkg/apis/kafka.strimzi.io/v1beta2/types.go";

    private static final ApiVersion API_VERSION = ApiVersion.V1BETA2;

    private static final List<String> IGNORED_PROPERTIES = List.of("apiVersion", "kind", "metadata");

    private static final List<Class<?>> CRDS = List.of(
            Kafka.class,
            KafkaNodePool.class,
            KafkaConnect.class,
            KafkaMirrorMaker2.class,
            KafkaBridge.class,
            KafkaRebalance.class,
            KafkaTopic.class,
            KafkaConnector.class,
            KafkaUser.class
    );

    private final OutputStreamWriter out;
    private final Stack<Class<?>> toBeGenerated = new Stack<>();
    private final Set<Class<?>> alreadyGenerated = new HashSet<>();

    public static void main(String[] args) throws IOException, InvocationTargetException, NoSuchMethodException, IllegalAccessException {
        LOGGER.info("Generating Strimzi Golang APIs into {}", OUTPUT_PATH);

        CodeGenerator codeGenerator = new CodeGenerator();
        codeGenerator.generate();
        codeGenerator.close();
    }

    private CodeGenerator() throws FileNotFoundException {
        out = new OutputStreamWriter(new FileOutputStream(OUTPUT_PATH), StandardCharsets.UTF_8);
    }

    private void generate() throws IOException, InvocationTargetException, NoSuchMethodException, IllegalAccessException {
        headerBoilerplate();
        generatePackage();
        generateImports();

        for (Class<?> crd : CRDS) {
            generateCrd(crd);
        }

        // Generate the next types
        while (!toBeGenerated.isEmpty()) {
            Class<?> propertyType = toBeGenerated.pop();
            alreadyGenerated.add(propertyType); // Add it to processed to not add it again
            generateType(propertyType);
        }
    }

    private void generateCrd(Class<?> crd) throws IOException {
        generateCrdList(crd);

        LOGGER.info("Generating {} CRD type", crd.getSimpleName());
        out.append(NL)
                .append("// +genclient").append(NL)
                .append("// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object").append(NL)
                .append(NL)
                .append("type ").append(crd.getSimpleName()).append(" struct {").append(NL)
                .append(TAB).append("metav1.TypeMeta `json:\",inline\"`").append(NL)
                .append(TAB).append("metav1.ObjectMeta `json:\"metadata,omitempty\"`").append(NL)
                .append(TAB).append(NL);

        Map<String, Property> properties = Property.properties(API_VERSION, crd);

        // Generate the CRD fields
        for (Property property : properties.values()) {
            if (!IGNORED_PROPERTIES.contains(property.getName())) {
                generateField(property, true);

                if (!toBeGenerated.contains(property.getType().getType()) && !alreadyGenerated.contains(property.getType().getType())) {
                    toBeGenerated.push(property.getType().getType());
                }
            }
        }

        out.append("}").append(NL);
    }

    private void generateField(Property property, boolean omitEmpty) throws IOException {
        PropertyType propertyType = property.getType();
        Class<?> returnType = propertyType.getType();

        if (returnType.getName().startsWith("io.fabric8.kubernetes.api.model.")) {
            generateFabric8Field(property.getGolangName(), returnType, property.getName());
        } else if (propertyType.getGenericType() instanceof ParameterizedType
                && ((ParameterizedType) propertyType.getGenericType()).getRawType().equals(Map.class)
                && ((ParameterizedType) propertyType.getGenericType()).getActualTypeArguments()[0].equals(Integer.class)) {
            LOGGER.error("Unsupported type {}", returnType);
        } else if (propertyType.getGenericType() instanceof ParameterizedType
                && ((ParameterizedType) propertyType.getGenericType()).getRawType().equals(Map.class)
                && propertyType.isMapOfTypes(String.class, Quantity.class)) {
            // Should not be needed as we replace Fabric8 types with Kube types instead of parsing them
            LOGGER.error("Unsupported type {}", returnType);
        } else if (Map.class.equals(returnType)) {
            if (propertyType.isMapOfTypes(String.class, String.class)) {
                generateField(property.getGolangName(), "map[string]string", property.getName(), omitEmpty);
            } else if (propertyType.isMapOfTypes(String.class, Object.class)) {
                generateField(property.getGolangName(), "MapStringObject", property.getName(), omitEmpty);
            } else {
                LOGGER.error("Unsupported Map type {}", returnType);
            }
        } else if (Utils.isJsonScalarType(returnType)) {
            if (returnType.isEnum()) {
                generateField(property.getGolangName(), returnType.getSimpleName(), property.getName(), omitEmpty);
                addToStackIfNeeded(returnType);
            } else {
                generateField(property.getGolangName(), typeName(returnType), property.getName(), omitEmpty);
            }
        } else if (returnType.isArray() || List.class.equals(returnType)) {
            generateArrayField(property, omitEmpty);
        } else {
            generateField(property.getGolangName(), "*" + property.getType().getType().getSimpleName(), property.getName(), omitEmpty);
            addToStackIfNeeded(property.getType().getType());
        }
    }

    private void addToStackIfNeeded(Class<?> type) {
        if (!toBeGenerated.contains(type) && !alreadyGenerated.contains(type)) {
            toBeGenerated.push(type);
        }
    }

    private String typeName(Class<?> type) {
        if (String.class.equals(type)) {
            return "string";
        } else if (byte.class.equals(type) || Byte.class.equals(type)) {
            return "int8";
        } else if (short.class.equals(type) || Short.class.equals(type)) {
            return "int16";
        } else if (int.class.equals(type) || Integer.class.equals(type)) {
            return "int32";
        } else if (long.class.equals(type) || Long.class.equals(type)) {
            return "int64";
        } else if (boolean.class.equals(type) || Boolean.class.equals(type)) {
            return "bool";
        } else if (float.class.equals(type) || Float.class.equals(type)) {
            return "float32";
        } else if (Double.class.equals(type) || double.class.equals(type)) {
            return "float64";
        } else {
            throw new RuntimeException(type.getName());
        }
    }

    private boolean isBasicType(Class<?> type) {
        return String.class.equals(type)
                || byte.class.equals(type)
                || Byte.class.equals(type)
                || short.class.equals(type)
                || Short.class.equals(type)
                || int.class.equals(type)
                || Integer.class.equals(type)
                || long.class.equals(type)
                || Long.class.equals(type)
                || boolean.class.equals(type)
                || Boolean.class.equals(type)
                || float.class.equals(type)
                || Float.class.equals(type)
                || Double.class.equals(type)
                || double.class.equals(type);
    }

    private void generateArrayField(Property property, boolean omitEmpty) throws IOException   {
        String arrayMarker = "[]".repeat(Math.max(0, property.getType().arrayDimension()));

        Class<?> elementType = property.getType().arrayBase();
        if (isBasicType(elementType)) {
            generateField(property.getGolangName(), arrayMarker + typeName(elementType), property.getName(), omitEmpty);
        } else if (elementType.getName().startsWith("io.fabric8.kubernetes.api.model.")) {
            generateField(property.getGolangName(), arrayMarker + mapFabric8TypeToKubernetes(elementType), property.getName(), omitEmpty);
        } else {
            generateField(property.getGolangName(), arrayMarker + elementType.getSimpleName(), property.getName(), omitEmpty);
            addToStackIfNeeded(elementType);
        }
    }

    private void generateFabric8Field(String goName, Class<?> type, String jsonName) throws IOException {
        out.append(TAB).append(goName).append(" *").append(mapFabric8TypeToKubernetes(type)).append(" ").append("`json:\"").append(jsonName).append(",omitempty\"`").append(NL);
    }

    private String mapFabric8TypeToKubernetes(Class<?> type) {
        return switch (type.getName()) {
            case "io.fabric8.kubernetes.api.model.Affinity" -> "corev1.Affinity";
            case "io.fabric8.kubernetes.api.model.CSIVolumeSource" -> "corev1.CSIVolumeSource";
            case "io.fabric8.kubernetes.api.model.ConfigMapKeySelector" -> "corev1.ConfigMapKeySelector";
            case "io.fabric8.kubernetes.api.model.ConfigMapVolumeSource" -> "corev1.ConfigMapVolumeSource";
            case "io.fabric8.kubernetes.api.model.EmptyDirVolumeSource" -> "corev1.EmptyDirVolumeSource";
            case "io.fabric8.kubernetes.api.model.HostAlias" -> "corev1.HostAlias";
            case "io.fabric8.kubernetes.api.model.LabelSelector" -> "corev1.LabelSelector";
            case "io.fabric8.kubernetes.api.model.LocalObjectReference" -> "corev1.LocalObjectReference";
            case "io.fabric8.kubernetes.api.model.PersistentVolumeClaimVolumeSource" -> "corev1.PersistentVolumeClaimVolumeSource";
            case "io.fabric8.kubernetes.api.model.PodDNSConfig" -> "corev1.PodDNSConfig";
            case "io.fabric8.kubernetes.api.model.PodSecurityContext" -> "corev1.PodSecurityContext";
            case "io.fabric8.kubernetes.api.model.PodSecurityContextBuilder" -> "corev1.PodSecurityContextBuilder";
            case "io.fabric8.kubernetes.api.model.ResourceRequirements" -> "corev1.ResourceRequirements";
            case "io.fabric8.kubernetes.api.model.SecretKeySelector" -> "corev1.SecretKeySelector";
            case "io.fabric8.kubernetes.api.model.SecretVolumeSource" -> "corev1.SecretVolumeSource";
            case "io.fabric8.kubernetes.api.model.SecurityContext" -> "corev1.SecurityContext";
            case "io.fabric8.kubernetes.api.model.Toleration" -> "corev1.Toleration";
            case "io.fabric8.kubernetes.api.model.TopologySpreadConstraint" -> "corev1.TopologySpreadConstraint";
            case "io.fabric8.kubernetes.api.model.VolumeMount" -> "corev1.VolumeMount";
            case "io.fabric8.kubernetes.api.model.networking.v1.NetworkPolicyPeer" -> "networkingv1.NetworkPolicyPeer";
            default -> {
                LOGGER.error("Unmapped Fabric8 class {}", type.getName());
                yield "";
            }
        };
    }

    private void generateField(String goName, String type, String jsonName, boolean omitEmpty) throws IOException {
        out.append(TAB).append(goName).append(" ").append(type).append(" ").append("`json:\"").append(jsonName).append(omitEmpty ? ",omitempty\"`" : "\"`").append(NL);
    }

    private void generateType(Class<?> type) throws IOException, NoSuchMethodException, InvocationTargetException, IllegalAccessException {
        if (type.isEnum()) {
            LOGGER.info("Generating {} enum type", type.getSimpleName());
            out.append(NL)
                    .append("type ").append(type.getSimpleName()).append(" string").append(NL)
                    .append("const (").append(NL);

            ObjectMapper objectMapper = new ObjectMapper();
            Method valuesMethod = type.getMethod("values");
            Enum<?>[] enums = (Enum<?>[]) valuesMethod.invoke(null);

            for (Enum<?> e : enums) {
                out.append(TAB).append(validConstName(e.toString(), type.getSimpleName())).append(" ").append(type.getSimpleName()).append(" = ").append(objectMapper.writeValueAsString(e)).append(NL);
            }

            out.append(")").append(NL);
        } else if (Utils.isPolymorphic(type)) {
            LOGGER.info("Generating {} poloymorphic type", type.getSimpleName());

            // Generates the type constant
            String typeType = type.getSimpleName() + "Type";
            List<String> typeNames = Property.subtypeNames(type);

            out.append(NL)
                    .append("type ").append(typeType).append(" string").append(NL)
                    .append("const (").append(NL);

            for (String typeName : typeNames) {
                out.append(TAB).append(validConstName(typeName, typeType)).append(" ").append(typeType).append(" = \"").append(typeName).append("\"").append(NL);
            }

            out.append(")").append(NL);

            // Generate union or properties
            List<Class<?>> subtypes = Property.subtypes(type);
            Map<String, Property> properties = new HashMap<>();
            for (Class<?> subtype : subtypes) {
                properties.putAll(Property.properties(API_VERSION, subtype));
            }

            out.append(NL)
                    .append("type ").append(type.getSimpleName()).append(" struct {").append(NL);

            for (Property property : properties.values()) {
                if ("type".equals(property.getName())) {
                    generateField(property.getGolangName(), typeType, property.getName(), true);
                } else {
                    if ("id".equals(property.getName())
                            && SingleVolumeStorage.class.getSimpleName().equals(type.getSimpleName())) {
                        // The id field for JBOD volumes should not be omitted when set to 0 when serializing the objects.
                        // So we need some special handling for it and not mark it as omit when empty.
                        // This does not seem to be detectable in any other way in the Strimzi Java classes, so we just hardcode it here.
                        LOGGER.info("Special handling for id field in SingleVolumeStorage");
                        generateField(property, false);
                    } else {
                        generateField(property, true);
                    }

                }
            }

            out.append("}").append(NL);
        } else {
            LOGGER.info("Generating {} type", type.getSimpleName());
            out.append(NL)
                    .append("type ").append(type.getSimpleName()).append(" struct {").append(NL);

            Map<String, Property> properties = Property.properties(API_VERSION, type);

            for (Property property : properties.values()) {
                if ("tls".equals(property.getName())
                        && GenericKafkaListener.class.getSimpleName().equals(type.getSimpleName())) {
                    // The tls field for listener configuration should not be omitted when set to false when serializing the objects.
                    // So we need some special handling for it and not mark it as omit when empty.
                    // This does not seem to be detectable in any other way in the Strimzi Java classes, so we just hardcode it here.
                    LOGGER.info("Special handling for tls field in GenericKafkaListener");
                    generateField(property, false);
                } else if ("maxUnavailable".equals(property.getName())
                        && PodDisruptionBudgetTemplate.class.getSimpleName().equals(type.getSimpleName())) {
                    // The maxUnavailable field for PodDisruptionBudgetTemplate should not be omitted when set to 0 when serializing the objects.
                    // So we need some special handling for it and not mark it as omit when empty.
                    // This does not seem to be detectable in any other way in the Strimzi Java classes, so we just hardcode it here.
                    LOGGER.info("Special handling for maxUnavailable field in PodDisruptionBudgetTemplate");
                    generateField(property, false);
                } else {
                    generateField(property, true);
                }
            }

            out.append("}").append(NL);
        }
    }

    private static String validConstName(String constName, String typeName) {
        return (constName + "_" + typeName).replace("-", "_").toUpperCase(Locale.ROOT);
    }

    private void generateCrdList(Class<?> crd) throws IOException {
        LOGGER.info("Generating {} CRD list type", crd.getSimpleName());

        out.append(NL)
                .append("// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object").append(NL)
                .append(NL)
                .append("type ").append(crd.getSimpleName()).append("List struct {").append(NL)
                .append(TAB).append("metav1.TypeMeta `json:\",inline\"`").append(NL)
                .append(TAB).append("metav1.ListMeta `json:\"metadata,omitempty\"`").append(NL)
                .append(TAB).append(NL)
                .append(TAB).append("Items []").append(crd.getSimpleName()).append(" `json:\"items\"`").append(NL)
                .append("}").append(NL);
    }

    private void close() throws IOException {
        out.flush();
        out.close();
    }

    private void headerBoilerplate() throws IOException {
        LOGGER.info("Adding boilerplate header from {}", HEADER_BOILERPLATE_PATH);
        out.append(Files.readString(HEADER_BOILERPLATE_PATH, StandardCharsets.UTF_8))
                .append(NL);
    }

    private void generatePackage() throws IOException {
        LOGGER.info("Setting package to {}", API_VERSION);
        out.append("package ").append(API_VERSION.toString()).append(NL);
    }

    private void generateImports() throws IOException {
        LOGGER.info("Generating imports");
        out.append(NL)
                .append("import (").append(NL)
                .append(TAB).append("corev1 \"k8s.io/api/core/v1\"").append(NL)
                .append(TAB).append("networkingv1 \"k8s.io/api/networking/v1\"").append(NL)
                .append(TAB).append("metav1 \"k8s.io/apimachinery/pkg/apis/meta/v1\"").append(NL)
                .append(")").append(NL);
    }
}
