package com.springboot.service.controller;

import com.springboot.service.annotation.TokenVerification;
import com.springboot.service.common.Lang.Result;
import com.springboot.service.entity.Groups;
import com.springboot.service.repository.GroupsRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.validation.annotation.Validated;
import org.springframework.web.bind.annotation.*;

@ResponseBody
@RestController
@RequestMapping("/api/v1/group")
public class GroupsController {

    @Autowired
    private GroupsRepository groupsRepository;

    /**
     * 创建权限
     * @param groups 创建的权限组
     * @return 返回创建的结果
     */
    @TokenVerification(group = TokenVerification.GROUP.ADMIN)
    @PutMapping("")
    public Result addGroup(@Validated @RequestBody Groups groups){
        if(groupsRepository.findAllByName(groups.getName())!=null){
            return Result.fail("权限信息已存在",groups);
        }
        Groups save = groupsRepository.save(groups);
        return Result.success("权限创建成功",save);
    }

    /**
     * 获取数据库中所有权限列表
     * @return 返回数组列表
     */
    @GetMapping()
    public Result getGroupList(){
        return Result.success(groupsRepository.findAll());
    }

    /**
     * 修改权限组的名字
     * @param id 需要修改的id
     * @param groups 修改后的姓名
     * @return 修改的结果
     */
    @TokenVerification(group = TokenVerification.GROUP.ADMIN)
    @PostMapping("/{id}")
    public Result setGroup(@PathVariable Integer id,@Validated @RequestBody Groups groups){
        if(!groupsRepository.findById(id).isPresent()){
            return Result.fail("id不存在");
        }
        Groups res = groupsRepository.save(new Groups(id, groups.getName()));
        return Result.success(res);
    }

    /**
     * 删除权限组
     * @param id 删除的id
     * @return 返回删除的结果
     */
    @TokenVerification(group = TokenVerification.GROUP.ADMIN)
    @DeleteMapping("/{id}")
    public Result deleteGroup(@PathVariable Integer id){
        if(!groupsRepository.findById(id).isPresent()){
            return Result.fail("权限组不存在");
        }
        groupsRepository.deleteById(id);
        return Result.success("删除成功");
    }
}
