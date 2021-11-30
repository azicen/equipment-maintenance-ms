package com.springboot.service.annotation;

import java.lang.annotation.ElementType;
import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;
import java.lang.annotation.Target;

@Target({ElementType.METHOD, ElementType.TYPE})
@Retention(RetentionPolicy.RUNTIME)
public @interface TokenVerification {
    enum GROUP {
        ADMIN(1),
        USER(0);
        private final int group;

        public int getGroup() {
            return group;
        }

        GROUP(int group) {
            this.group = group;
        }
    }
    GROUP group() default GROUP.USER;
}
