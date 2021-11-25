package com.springboot.service.controller;

import cn.hutool.core.map.MapBuilder;
import cn.hutool.core.map.MapUtil;
import com.springboot.service.common.lang.Result;
import com.springboot.service.entity.User;
import com.springboot.service.Repository.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@ResponseBody
@RestController
@RequestMapping("/user")
public class UserController {
    @Autowired
    private UserRepository userRepository;

    @GetMapping("/findAll")
    public List<User> findAll(){
        return userRepository.findAll();
    }

    @RequestMapping("/{id}")
    public Result getUser(@PathVariable("id") Integer id){
        User user = userRepository.getOne(id);
        System.out.println(Result.success(
                MapUtil.builder().put("name",user.getName())
                        .put("status",user.getStatus())
                        .put("group",user.getId())));
        return
                Result.success(
                MapUtil.builder().put("name",user.getName())
                .put("status",user.getStatus())
                .put("group",user.getId()).map());

    }
}
