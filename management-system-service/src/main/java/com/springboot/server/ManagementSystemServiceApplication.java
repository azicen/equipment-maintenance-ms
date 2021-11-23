package com.springboot.server;

import org.mybatis.spring.annotation.MapperScan;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
@MapperScan("com.springboot.server.mapper")
public class ManagementSystemServiceApplication {
    public static void main(String[] args) {
        SpringApplication.run(ManagementSystemServiceApplication.class, args);

    }

}
