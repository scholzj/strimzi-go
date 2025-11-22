package cz.scholz.strimzi.golang.generator;

import io.strimzi.api.annotations.ApiVersion;
import io.strimzi.api.kafka.model.bridge.KafkaBridge;
import io.strimzi.api.kafka.model.connect.KafkaConnect;
import io.strimzi.api.kafka.model.connector.KafkaConnector;
import io.strimzi.api.kafka.model.kafka.Kafka;
import io.strimzi.api.kafka.model.mirrormaker2.KafkaMirrorMaker2;
import io.strimzi.api.kafka.model.nodepool.KafkaNodePool;
import io.strimzi.api.kafka.model.podset.StrimziPodSet;
import io.strimzi.api.kafka.model.rebalance.KafkaRebalance;
import io.strimzi.api.kafka.model.topic.KafkaTopic;
import io.strimzi.api.kafka.model.user.KafkaUser;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;

import java.io.IOException;
import java.lang.reflect.InvocationTargetException;
import java.nio.file.Path;
import java.util.List;

public class Main {
    private final static Logger LOGGER = LogManager.getLogger(Main.class);

    private static final Path HEADER_BOILERPLATE_PATH = Path.of("../hack/boilerplate.go.txt");

    private static final String KAFKA_V1BETA2_OUTPUT_PATH = "../pkg/apis/kafka.strimzi.io/v1beta2/types.go";
    private static final String KAFKA_V1_OUTPUT_PATH = "../pkg/apis/kafka.strimzi.io/v1/types.go";
    private static final List<Class<?>> KAFKA_CRDS = List.of(
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

    private static final String CORE_V1BETA2_OUTPUT_PATH = "../pkg/apis/core.strimzi.io/v1beta2/types.go";
    private static final String CORE_V1_OUTPUT_PATH = "../pkg/apis/core.strimzi.io/v1/types.go";
    private static final List<Class<?>> CORE_CRDS = List.of(
            StrimziPodSet.class
    );

    public static void main(String[] args) throws IOException, InvocationTargetException, NoSuchMethodException, IllegalAccessException {
        // Generate v1beta2 versions
        LOGGER.info("Generating Strimzi Golang v1beta2 APIs into {}", KAFKA_V1BETA2_OUTPUT_PATH);

        CodeGenerator codeGenerator = new CodeGenerator(KAFKA_CRDS, List.of("corev1", "networkingv1", "metav1"), ApiVersion.V1BETA2, HEADER_BOILERPLATE_PATH, KAFKA_V1BETA2_OUTPUT_PATH);
        codeGenerator.generate();
        codeGenerator.close();

        codeGenerator = new CodeGenerator(CORE_CRDS, List.of("kafkav1beta2", "metav1"), ApiVersion.V1BETA2, HEADER_BOILERPLATE_PATH, CORE_V1BETA2_OUTPUT_PATH);
        codeGenerator.generate();
        codeGenerator.close();

        // Generate v1 versions
        LOGGER.info("Generating Strimzi Golang v1 APIs into {}", KAFKA_V1_OUTPUT_PATH);

        codeGenerator = new CodeGenerator(KAFKA_CRDS, List.of("corev1", "networkingv1", "metav1"), ApiVersion.V1, HEADER_BOILERPLATE_PATH, KAFKA_V1_OUTPUT_PATH);
        codeGenerator.generate();
        codeGenerator.close();

        codeGenerator = new CodeGenerator(CORE_CRDS, List.of("kafkav1", "metav1"), ApiVersion.V1, HEADER_BOILERPLATE_PATH, CORE_V1_OUTPUT_PATH);
        codeGenerator.generate();
        codeGenerator.close();
    }
}
