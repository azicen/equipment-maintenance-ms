package com.springboot.service.Controller;

import cn.hutool.core.map.MapUtil;
import com.springboot.service.Entity.Groups;
import com.springboot.service.Entity.User;
import com.springboot.service.Entity.UserInGroup;
import com.springboot.service.Repository.GroupsRepository;
import com.springboot.service.Repository.UserInGroupRepository;
import com.springboot.service.Repository.UserRepository;
import com.springboot.service.common.Lang.Result;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.domain.PageRequest;
import org.springframework.data.domain.Pageable;
import org.springframework.web.bind.annotation.*;

import java.util.ArrayList;
import java.util.List;
import java.util.Optional;

@ResponseBody
@RestController
@RequestMapping("/api/v1/user")
public class UserController {
    /**
     * 用户表
     */
    @Autowired
    private UserRepository userRepository;
    /**
     * 用户对应权限关系表
     */
    @Autowired
    private UserInGroupRepository userInGroupRepository;
    /**
     * 权限表
     */
    @Autowired
    private GroupsRepository groupsRepository;

    /**
     * 添加用户
     * @param name 用户姓名
     * @param passwd 用户密码
     * @return 会根据数据库中是否存在对应name的用户返回是否创建成功
     */
    @PutMapping("/")
    public Result addUser(@RequestHeader("name") String name, @RequestHeader("passwd") String passwd){
        User user=userRepository.save(new User(null,name,passwd,false));
        return Result.success(MapUtil.builder().put("id",user.getId()).map());
    }

    /**
     * 将用户添加到权限组
     * @param userId 用户id
     * @param groupId 权限组id
     * @return 返回是否添加成功
     */
    @PutMapping("/{id}/group/{group}")
    public Result addUserGroup(@PathVariable("id") Integer userId,@PathVariable("group") Integer groupId){
        if(userInGroupRepository.findAllByUserIdAndGroupId(userId,groupId)!=null){
            return Result.fail("用户所对应的权限组已存在");
        }
        if(!userRepository.findById(userId).isPresent()){
            return Result.fail("用户不存在");
        }
        if(!groupsRepository.findById(groupId).isPresent()){
            return Result.fail("权限组不存在");
        }
        UserInGroup uIg = userInGroupRepository.save(new UserInGroup(null,userId,groupId));
        return Result.success("权限添加成功",MapUtil.builder()
                .put("id",uIg.getUserId())
                .put("group",uIg.getGroupId()).map());
    }

    /**
     * 删除指定用户的权限组
     * @param id 用户id
     * @param groupId 权限组id
     * @return 返回是否删除
     */
    @DeleteMapping("/{id}/group/{group}")
    public Result deleteUserInGroup(@PathVariable("id") Integer id,@PathVariable("group") Integer groupId){
        UserInGroup uIg = userInGroupRepository.findAllByUserIdAndGroupId(id, groupId);
        if(uIg==null){
            return Result.fail("用户不在此权限组或用户不存在");
        }
        userInGroupRepository.delete(uIg);
        return Result.success("删除成功");
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
        userRepository.deleteAllById(id);
        userInGroupRepository.removeAllByUserId(id);
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
            ,@RequestHeader("passwd") String passwd,@RequestHeader("status")Boolean status){
        Optional<User> option=userRepository.findById(id);
        if(!option.isPresent()){
            return Result.fail(Result.CODE.NOT_FOND,"未找到该用户",null);
        }
        User user=option.get();
        user.setName(name);
        user.setPasswd(passwd);
        user.setStatus(status);
        userRepository.save(user);
        return Result.success(MapUtil.builder().put("name",user.getName())
                .put("status",user.getStatus()).map());
    }

    /**
     * 返回用户权限组
     * @param id 需要访问的用户id
     * @return 返回用户所在的所有权限组
     */
    @GetMapping("/{id}/group")
    public List<Object> getUserGroup(@PathVariable("id") Integer id){
        List<UserInGroup> g=userInGroupRepository.findAllByUserId(id);
        List<Object> userInGroup=new ArrayList<>();
        g.forEach((UserInGroup u)->{
            userInGroup.add(MapUtil.builder().put("group",u.getGroupId()).map());
        });
        return userInGroup;
    }

    /**
     * 根据id查找用户
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
                        .put("status",user.getStatus()).map());
    }

    /**
     * 分页查询
     * @param n 起始页数
     * @param page 当前页数有多少条数据
     * @return 返回object列表表示当前页数的数据对象
     */
    @GetMapping("/list")
    public List<Object> getBeginToEndUser(int n, int page){
        Pageable pageable= PageRequest.of(page,n);
        List<User> findRes=userRepository.findAll(pageable).getContent();
        List<Object> res=new ArrayList<>();
        findRes.forEach((User u)-> res.add(MapUtil.builder()
                .put("status",u.getStatus())
                .put("name",u.getName())
                .put("id",u.getId())
                .map()));
        return res;
    }
}
