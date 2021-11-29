package com.springboot.service.Repository;

import com.springboot.service.Entity.UserInGroup;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.transaction.annotation.Transactional;


public interface UserInGroupRepository extends JpaRepository<UserInGroup,Integer> {
    @Transactional
    void removeAllByUserId(Integer id);
}
