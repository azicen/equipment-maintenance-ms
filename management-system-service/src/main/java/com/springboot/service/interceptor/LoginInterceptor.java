package com.springboot.service.interceptor;

import cn.hutool.core.map.MapUtil;
import com.alibaba.fastjson.JSON;
import com.auth0.jwt.interfaces.DecodedJWT;
import com.springboot.service.annotation.TokenVerification;
import com.springboot.service.common.Lang.Result;
import com.springboot.service.utils.JWTUtils;
import org.springframework.stereotype.Component;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.method.HandlerMethod;
import org.springframework.web.servlet.HandlerInterceptor;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;
import java.io.ObjectOutputStream;
import java.lang.reflect.Method;

@Component
public class LoginInterceptor implements HandlerInterceptor {

    @Override
    public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object handler) throws Exception {
        response.setCharacterEncoding("UTF-8");
        response.setContentType("application/json");
        //不是映射到方法通过
        if(!(handler instanceof HandlerMethod)){
            response.getWriter().write(JSON.toJSONString(Result.fail("未找到api")));
            return false;
        }
        //映射到方法且有TokenPass注解
        Method method=((HandlerMethod)handler).getMethod();
        //判断是否有token验证的注解，用于区分权限组并拦截对应的api访问
        if(method.isAnnotationPresent(TokenVerification.class)){
            TokenVerification tokenPass=method.getAnnotation(TokenVerification.class);
            String token=request.getHeader("Authorization");
            //如果没有token拦截
            if(token==null){
                response.getWriter().write(JSON.toJSONString(Result.fail("访问此api需要token信息")));
                return false;
            }
            Result tokenV = JWTUtils.verify(token);
            if(tokenV.getCode()==200){
                //解析token对应的group字段判断权限组是否达到访问该方法需要的级别
                String group = ((DecodedJWT) tokenV.getData()).getClaim("group").asString();
                boolean isGroup=Integer.parseInt(group)>=tokenPass.group().getGroup();
                if(!isGroup){
                    response.getWriter().write(JSON.toJSONString(Result.fail("权限等级不够")));
                }
                return isGroup;
            }else{
                //token认证失败拦截
                response.getWriter().write(JSON.toJSONString(Result.fail("token认证失败")));
                return false;
            }
        }
        return true;
    }
}
