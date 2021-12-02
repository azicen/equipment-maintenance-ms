package com.springboot.service.controller;

import com.springboot.service.annotation.TokenVerification;
import com.springboot.service.common.Lang.Result;
import com.springboot.service.entity.Equipment;
import com.springboot.service.repository.EquipmentRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.validation.annotation.Validated;
import org.springframework.web.bind.annotation.*;

import java.sql.Timestamp;
import java.util.Optional;

@ResponseBody
@RestController
@RequestMapping("/api/v1/equipment")
public class EquipmentController {
    @Autowired
    private EquipmentRepository equipmentRepository;

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


}
