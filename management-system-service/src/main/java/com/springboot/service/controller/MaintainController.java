package com.springboot.service.controller;

import com.springboot.service.common.Lang.Result;
import com.springboot.service.entity.Maintain;
import com.springboot.service.repository.MaintainRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.validation.annotation.Validated;
import org.springframework.web.bind.annotation.*;

@ResponseBody
@RestController
@RequestMapping("/api/v1/Maintain")
public class MaintainController {

    @Autowired
    private MaintainRepository maintainRepository;


    /**
     * 获取数据列表
     * @return 返回数据列表
     */
    @GetMapping("")
    public Result getList(){
        return Result.success(maintainRepository.findAll());
    }

    /**
     * 添加数据
     * @param m 添加数据
     * @return 添加数据的结果
     */
    @PutMapping("")
    public Result addList(@Validated @RequestBody Maintain m){
        return Result.success(maintainRepository.save(m));
    }
}
