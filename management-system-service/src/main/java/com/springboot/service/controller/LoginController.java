package com.springboot.service.controller;

import com.springboot.service.entity.User;
import com.springboot.service.entity.UserInGroup;
import com.springboot.service.repository.UserInGroupRepository;
import com.springboot.service.repository.UserRepository;
import com.springboot.service.common.Lang.Result;
import com.springboot.service.utils.JWTUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.validation.annotation.Validated;
import org.springframework.web.bind.annotation.*;

import javax.servlet.http.HttpServletResponse;
import java.util.HashMap;

@ResponseBody
@RestController
@RequestMapping("/api/v1/login")
public class LoginController {
    @Autowired
    private UserRepository userRepository;
    @Autowired
    private UserInGroupRepository userInGroupRepository;

    @PostMapping("")
    public Result userLogin(@Validated @RequestBody User u, HttpServletResponse response){
        User user = userRepository.findAllByIdAndPasswd(u.getId(), u.getPasswd());
        if(user==null){
            return Result.fail("用户不存在或密码错误",u);
        }
        UserInGroup uIg=userInGroupRepository.findAllByUserIdAndGroupId(user.getId(),1);
//        System.out.println(uIg==null?"0":"1");
        HashMap<String, String> map = new HashMap<>();
        map.put("id",user.getId().toString());
        map.put("group",uIg==null?"0":"1");
        String token= JWTUtils.getToken(map);
        response.setHeader("Authorization",token);
        response.setHeader("Access-control-expose-Headers", "Authorization");
        return Result.success("登录成功");
    }
}
