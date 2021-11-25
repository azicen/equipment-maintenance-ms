package com.springboot.service.entity;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import lombok.*;
import javax.persistence.*;

@Entity
@Data
@NoArgsConstructor
@AllArgsConstructor
@Table(name = "tb_user")
@JsonIgnoreProperties(value={"hibernateLazyInitializer","handler","fieldHandler"})
public class User {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Integer id;         //  id
    @Column(length = 20)
    private String name;        //姓名
    @Column(length = 20)
    private String passWord;    //密码
    @Column(length = 20)
    private String status;     //账号状态
}
