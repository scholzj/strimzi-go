package cz.scholz.strimzi.golang.generator;

import com.fasterxml.jackson.annotation.JsonSubTypes;

import java.lang.reflect.Method;

public class Utils {
    static boolean isBoxedPrimitive(Class<?> cls) {
        boolean intLike = Short.class.equals(cls)
                || Integer.class.equals(cls)
                || Long.class.equals(cls);
        return Boolean.class.equals(cls)
                || intLike
                || Float.class.equals(cls)
                || Double.class.equals(cls)
                || Character.class.equals(cls);
    }

    static boolean isJsonScalarType(Class<?> cls) {
        return cls.isPrimitive()
                || isBoxedPrimitive(cls)
                || cls.equals(String.class)
                || cls.isEnum();
    }

    static boolean isPolymorphic(Class<?> cls) {
        return cls.isAnnotationPresent(JsonSubTypes.class);
    }

    static boolean isGetterName(Method method) {
        String name = method.getName();
        return name.startsWith("get")
                && name.length() > 3
                && isReallyGetterName(method, name)
                || name.startsWith("is")
                && name.length() > 2
                && method.getReturnType().equals(boolean.class);
    }

    private static boolean isReallyGetterName(Method method, String name) {
        return !"getClass".equals(name)
                && !("getDeclaringClass".equals(name) && Enum.class.equals(method.getDeclaringClass()));
    }
}
