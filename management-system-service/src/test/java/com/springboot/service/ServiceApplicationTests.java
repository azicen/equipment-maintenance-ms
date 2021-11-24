package com.springboot.service;

import com.springboot.service.mapper.UserMapper;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

@SpringBootTest
class ServiceApplicationTests {
    @Autowired
    UserMapper userMapper;

    @Test
    void contextLoads() {
        System.out.println(userMapper.findAll());
    }

}
