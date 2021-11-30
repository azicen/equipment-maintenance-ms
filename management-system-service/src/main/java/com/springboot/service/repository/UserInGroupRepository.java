package com.springboot.service.repository;

import com.springboot.service.entity.UserInGroup;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;


public interface UserInGroupRepository extends JpaRepository<UserInGroup,Integer> {
    @Transactional
    void removeAllByUserId(Integer id);
    UserInGroup findAllByUserIdAndGroupId(Integer userId,Integer groupId);
    List<UserInGroup> findAllByUserId(Integer id);
}
