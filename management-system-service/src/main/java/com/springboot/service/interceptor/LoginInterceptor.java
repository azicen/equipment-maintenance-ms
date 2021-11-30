package com.springboot.service.interceptor;

import com.auth0.jwt.interfaces.DecodedJWT;
import com.springboot.service.common.Lang.Result;
import com.springboot.service.utils.JWTUtils;
import org.springframework.stereotype.Component;
import org.springframework.web.servlet.HandlerInterceptor;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

@Component
public class LoginInterceptor implements HandlerInterceptor {
    @Override
    public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object handler) throws Exception {
        if(request.getRequestURI().equals("/login/")){
            return true;
        }
        String token=request.getHeader("Authorization");
        if(token==null){
            System.out.println("token不存在");
            return false;
        }
        Result tokenV = JWTUtils.verify(token);
        if(tokenV.getCode()==200){
//            if(((DecodedJWT)tokenV.getData()).getClaim("group").asBoolean())
            return true;
        }
        return false;
    }
}
