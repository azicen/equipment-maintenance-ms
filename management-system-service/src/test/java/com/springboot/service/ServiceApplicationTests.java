package com.springboot.service;

import com.springboot.service.entity.User;
import com.springboot.service.mapper.userMapper;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

import java.util.List;

@SpringBootTest
class ServiceApplicationTests {
    @Autowired
    userMapper userMapper;

    @Test
    void contextLoads() {
        System.out.println(userMapper.findAll());
    }

}
