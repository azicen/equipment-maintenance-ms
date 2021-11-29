package com.springboot.service;

import cn.hutool.crypto.asymmetric.Sign;
import com.auth0.jwt.JWT;
import com.auth0.jwt.JWTVerifier;
import com.auth0.jwt.algorithms.Algorithm;
import com.auth0.jwt.interfaces.DecodedJWT;
import com.auth0.jwt.interfaces.Verification;
import com.springboot.service.Controller.UserController;
import com.springboot.service.Repository.UserRepository;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

import javax.xml.crypto.Data;
import java.util.Calendar;
import java.util.HashMap;

@SpringBootTest
class ServiceApplicationTests {

    @Test
    void contextLoads() {
        Calendar i=Calendar.getInstance();
        i.add(Calendar.SECOND,50);
        HashMap<String ,Object> map=new HashMap<>();
        String token = JWT.create().withHeader(map)
                .withClaim("userId", 21)
                .withClaim("username", "aa")
                .withExpiresAt(i.getTime())
                .sign(Algorithm.HMAC256("aaa"));
        System.out.println(token);
        JWTVerifier ver = JWT.require(Algorithm.HMAC256("aaa")).build();
        DecodedJWT verify = ver.verify(token);
        System.out.println(verify.getClaim("userId").asInt());
    }

}
