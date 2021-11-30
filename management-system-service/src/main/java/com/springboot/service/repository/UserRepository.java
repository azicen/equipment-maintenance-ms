package com.springboot.service.repository;

import com.springboot.service.entity.User;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.transaction.annotation.Transactional;


public interface UserRepository extends JpaRepository<User,Integer> {
    @Transactional
    void deleteAllById(Integer id);
    User findAllByIdAndPasswd(Integer id,String passwd);
}
