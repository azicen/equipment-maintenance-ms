package com.springboot.service.controller;

import cn.hutool.core.map.MapUtil;
import com.springboot.service.common.lang.Result;
import com.springboot.service.entity.User;
import com.springboot.service.Repository.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.Optional;

@ResponseBody
@RestController
@RequestMapping("/api/v1/user")
public class UserController {
    @Autowired
    private UserRepository userRepository;

    @GetMapping("/findAll")
    public List<User> findAll(){
        return userRepository.findAll();
    }

    /**
     * get根据查找用户
     * @param id 用户id
     * @return 返回用户的名字,状态码,和用户组
     */
    @GetMapping("/{id}")
    public Result getUser(@PathVariable("id") Integer id){
        Optional<User> option=userRepository.findById(id);
        if(!option.isPresent()){
            return Result.fail(Result.CODE.NOT_FOND,"未找到该用户",null);
        }
        User user=option.get();
        return Result.success("查找成功",
                MapUtil.builder().put("name", user.getName())
                .put("status",user.getStatus())
                .put("group",user.getId()).map());
    }

    /**
     * 添加用户
     * @param name 用户姓名
     * @param passwd 用户密码
     * @return 会根据数据库中是否存在对应name的用户返回是否创建成功
     */
    @PutMapping("/")
    public Result addUser(@RequestHeader("name") String name,@RequestHeader("passwd") String passwd){
        List<User> userList=userRepository.findByName(name);
        if(!userList.isEmpty()){
            return Result.fail("用户已存在");
        }
        User user=userRepository.save(new User(null,name,passwd,"0"));
        return Result.success(MapUtil.builder().put("name",user.getName())
                .put("id",user.getId()).map());
    }

    /**
     * 根据id删除用户
     * @param id 要删除的用户id
     * @return 返回是否成功删除
     */
    @DeleteMapping("/{id}")
    public Result deleteUser(@PathVariable("id") Integer id){
        if(!userRepository.findById(id).isPresent()){
            return Result.fail(Result.CODE.NOT_FOND,"未找到该用户",null);
        }
        userRepository.removeAllById(id);
        return Result.success("删除成功");
    }

    /**
     *  修改用户数据
     * @param id 需要修改的用户id,由路由/api/user/id获得
     * @param name 修改后的姓名
     * @param passwd 修改后的密码
     * @param status 修改后的状态码
     * @return 根据是否存在用户返回对应的结果,若name重复则不能修改
     */
    @PostMapping("/{id}")
    public Result updateUser(@PathVariable("id") Integer id,@RequestHeader("name") String name
            ,@RequestHeader("passwd") String passwd,@RequestHeader("status")String status){
        Optional<User> option=userRepository.findById(id);
        if(!option.isPresent()){
            return Result.fail(Result.CODE.NOT_FOND,"未找到该用户",null);
        }
        if(!userRepository.findByName(name).isEmpty()){
            return Result.fail("用户名已存在,请重新设置用户名");
        }
        User user=option.get();
        user.setName(name);
        user.setPassWord(passwd);
        user.setStatus(status);
        userRepository.save(user);
        return Result.success(MapUtil.builder().put("name",user.getName())
                        .put("status",user.getStatus()).map());
    }


}
