package com.springboot.server;

import com.springboot.server.entity.User;
import com.springboot.server.entity.UserMapper;
import org.mybatis.spring.annotation.MapperScan;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

import javax.annotation.Resource;
import java.util.List;

@SpringBootApplication
@MapperScan("com.springboot.server.entity")
public class ManagementSystemServiceApplication {


    public static void main(String[] args) {
        SpringApplication.run(ManagementSystemServiceApplication.class, args);

    }

}
