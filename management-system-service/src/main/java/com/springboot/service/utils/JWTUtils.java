package com.springboot.service.utils;

import com.auth0.jwt.JWT;
import com.auth0.jwt.JWTCreator;
import com.auth0.jwt.algorithms.Algorithm;
import com.auth0.jwt.interfaces.DecodedJWT;
import com.springboot.service.common.Lang.Result;

import java.util.Calendar;
import java.util.Map;

public class JWTUtils {

    private static final String SING="sbSpringBoot";

    /**
     * 生成token
     * @param map 属性的键值对
     * @return 生成的token
     */
    public static String getToken(Map<String,String > map){
        Calendar instance =Calendar.getInstance();
        instance.add(Calendar.DATE,1);
        JWTCreator.Builder builder = JWT.create();
        map.forEach(builder::withClaim);
        return builder.withExpiresAt(instance.getTime()).sign(Algorithm.HMAC256(SING));
    }

    /**
     * token认证登录
     * @param token 需要认证的token
     * @return 认证的结果
     */
    public static Result verify(String token){
        DecodedJWT v = null;
        try {
            v = JWT.require(Algorithm.HMAC256(SING)).build().verify(token);
        }catch (Exception e){
            return Result.fail("token认证失败");
        }
        return Result.success("认证成功",v);
    }
}
