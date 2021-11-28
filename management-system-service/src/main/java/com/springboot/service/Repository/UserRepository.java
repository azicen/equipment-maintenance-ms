package com.springboot.service.Repository;

import com.springboot.service.Entity.User;
import org.springframework.data.jpa.repository.JpaRepository;

public interface UserRepository extends JpaRepository<User,Integer> {
//    List<User> findByName(String name);
//    Integer removeAllById(Integer id);
}
