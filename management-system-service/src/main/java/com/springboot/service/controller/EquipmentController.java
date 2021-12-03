package com.springboot.service.controller;

import com.springboot.service.annotation.TokenVerification;
import com.springboot.service.common.Lang.Result;
import com.springboot.service.entity.Equipment;
import com.springboot.service.entity.EquipmentInGroup;
import com.springboot.service.repository.EquipmentInGroupRepository;
import com.springboot.service.repository.EquipmentRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.domain.PageRequest;
import org.springframework.data.domain.Pageable;
import org.springframework.validation.annotation.Validated;
import org.springframework.web.bind.annotation.*;

import java.sql.Timestamp;
import java.util.List;
import java.util.Optional;

@ResponseBody
@RestController
@RequestMapping("/api/v1/equipment")
public class EquipmentController {
    @Autowired
    private EquipmentRepository equipmentRepository;

    @Autowired
    private EquipmentInGroupRepository equipmentInGroupRepository;

    /**
     * 创建设备
     *
     * @param equipment 创建的设备
     * @return 返回创建的结果
     */
    @TokenVerification(group = TokenVerification.GROUP.ADMIN)
    @PutMapping("")
    public Result addEquipment(@Validated @RequestBody Equipment equipment) {
        equipment.setCreateDate(new Timestamp(System.currentTimeMillis()).getTime());
        System.out.println(equipment);
        return Result.success(equipmentRepository.save(equipment));
    }

    /**
     * {
     *     "name":"洗衣机",
     *     "location":"宿舍楼",
     *     "status":"损坏",
     *     "startDate":78126,
     *     "deadLine":932,
     *     "typeId":"1",
     *     "userId":1,
     *     "createdate":""
     * }
     */

    /**
     * 获取设备信息
     *
     * @param id 设备的id号
     * @return 返回设备的信息
     */
    @GetMapping("/{id}")
    public Result getEquipment(@PathVariable Integer id) {
        Optional<Equipment> equipment = equipmentRepository.findById(id);
        return equipment.map(Result::success).orElseGet(() -> Result.fail("未找到对应id的数据"));
    }

    /**
     * 修改对应id的设备
     * @param id 需要修改的设备id
     * @param equipment 修改后的设备id
     * @return 返回修改的结果
     */
    @TokenVerification(group = TokenVerification.GROUP.ADMIN)
    @PostMapping("/{id}")
    public Result setEquipment(@PathVariable Integer id, @Validated @RequestBody Equipment equipment) {
        Optional<Equipment> findEquipment = equipmentRepository.findById(id);
        return findEquipment.map((e) -> {
            e.setName(equipment.getName());
            e.setStatus(equipment.getStatus());
            e.setLocation(equipment.getLocation());
            return Result.success("修改成功", e);
        }).orElseGet(() -> Result.fail("修改失败"));
    }

    /**
     * 删除设备
     * @param id 需要删除的设备id
     * @return 返回删除的结果
     */
    @TokenVerification(group = TokenVerification.GROUP.ADMIN)
    @DeleteMapping("/{id}")
    public Result deleteEquipment(@PathVariable Integer id){
        Optional<Equipment> equipment=equipmentRepository.findById(id);
        return equipment.map((e)->{
            equipmentRepository.deleteById(id);
            return Result.success("删除成功");
        }).orElseGet(()->Result.fail("未找到对应id的设备"));
    }

    /**
     * 分页查找
     * @param n 每页的个数
     * @param page 第几页
     * @return 返回查询的结果
     */
    @GetMapping("/list")
    public Result getEquipmentPage(int n,int page){
        Pageable pageable= PageRequest.of(page,n);
        List<Equipment> findRes=equipmentRepository.findAll(pageable).getContent();
        return Result.success(findRes);
    }

    /**
     * 将设备添加到对应分组
     * @param id 设备id
     * @param group 分组id
     * @return 返回添加的结果
     */
    @TokenVerification(group = TokenVerification.GROUP.ADMIN)
    @PutMapping("/{id}/group/{group}")
    public Result addEquipmentGroup(@PathVariable Integer id,@PathVariable Integer group){
        if(equipmentInGroupRepository.findAllByEquipmentIdAndGroupId(id, group)!=null){
            return Result.fail("对应权限组已存在");
        }
        return Result.success(equipmentInGroupRepository.save(new EquipmentInGroup(null,id,group)));
    }

    /**
     * 获取对应设备id的设备组
     * @param id 设备id
     * @return 返回查找结果
     */
    @GetMapping("/{id}/group")
    public Result getEquipmentGroup(@PathVariable Integer id){
        List<EquipmentInGroup> res = equipmentInGroupRepository.findAllByEquipmentId(id);
        return res==null?Result.fail("未找到对应数据"):Result.success(res);
    }

    /**
     * 删除对应设备的分组
     * @param id 设备id
     * @param groupId 分组id
     * @return 删除结果
     */
    @TokenVerification(group = TokenVerification.GROUP.ADMIN)
    @DeleteMapping("/{id}/group/{group}")
    public Result deleteEquipmentGroup(@PathVariable("id")Integer id,@PathVariable("group") Integer groupId){
        EquipmentInGroup res = equipmentInGroupRepository.findAllByEquipmentIdAndGroupId(id, groupId);
        if(res==null){
            return Result.fail("未找到对应数据");
        }
        equipmentInGroupRepository.delete(res);
        return Result.success("删除成功",res);
    }

}
