package com.springboot.service;

import com.springboot.service.Controller.UserController;
import com.springboot.service.Repository.UserRepository;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

@SpringBootTest
class ServiceApplicationTests {
    @Autowired
    UserRepository userMapper;
    @Autowired
    UserController userController;
    @Test
    void contextLoads() {

    }

}
