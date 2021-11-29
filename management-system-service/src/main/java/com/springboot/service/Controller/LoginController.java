package com.springboot.service.Controller;

import cn.hutool.core.map.MapUtil;
import com.springboot.service.Entity.User;
import com.springboot.service.Repository.UserRepository;
import com.springboot.service.common.Lang.Result;
import com.springboot.service.utils.JWTUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import javax.servlet.http.HttpServletResponse;
import java.util.HashMap;

@ResponseBody
@RestController
@RequestMapping("/login")
public class LoginController {
    @Autowired
    private UserRepository userRepository;

    @PostMapping("/")
    public Result userLogin(@RequestHeader("id") Integer id, @RequestHeader("passwd") String passwd
            , HttpServletResponse response){
        User user = userRepository.findAllByIdAndPasswd(id, passwd);
        if(user==null){
            return Result.fail("用户不存在或密码错误");
        }
        HashMap<String, String> map = new HashMap<>();
        map.put("id",user.getId().toString());
        map.put("passwd",user.getPasswd());
        String token= JWTUtils.getToken(map);
        response.setHeader("Authorization",token);
        response.setHeader("Access-control-expose-Headers", "Authorization");
        System.out.println(token);
        return Result.success("登录成功");
    }
}
