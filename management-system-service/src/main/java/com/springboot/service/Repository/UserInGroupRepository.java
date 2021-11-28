package com.springboot.service.Repository;

import com.springboot.service.Entity.UserInGroup;
import org.springframework.data.jpa.repository.JpaRepository;

public interface UserInGroupRepository extends JpaRepository<UserInGroup,Integer> {
}
