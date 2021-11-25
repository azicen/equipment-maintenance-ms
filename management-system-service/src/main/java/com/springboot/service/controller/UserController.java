package com.springboot.service.controller;

import cn.hutool.core.map.MapUtil;
import com.springboot.service.common.lang.Result;
import com.springboot.service.entity.User;
import com.springboot.service.Repository.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.Optional;

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
        Optional<User> option=userRepository.findById(id);
        if(!option.isPresent()){
            return Result.fail(Result.CODE.NOT_FOND,"未找到该用户",null);
        }
        User user=option.get();
        return Result.success(
                MapUtil.builder().put("name",user.getName())
                .put("status",user.getStatus())
                .put("group",user.getId()).map());

    }
}
