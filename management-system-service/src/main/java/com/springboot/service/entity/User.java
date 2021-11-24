package com.springboot.service.entity;

import lombok.*;
import org.hibernate.Hibernate;

import javax.persistence.*;
import java.util.Objects;

@Entity
@Data
@Table(name = "tb_user")
public class User {
    @Id
    @GeneratedValue
    private Integer id;
    @Column(length = 20)
    private String stuName;
    private int age;
}
