# go-bypass
go语言免杀项目。
声明：该项目来自作者日常学习笔记。 请勿利用相关技术以及工具从事非法测试，如因此产生的一切不良后果作者无关。
## 使用方式：
1、用bagua_en对CS生成的shellcode进行加密
![image](https://user-images.githubusercontent.com/94209165/196105104-b3c7f3d2-d341-43bd-93fb-bb9d25741f2b.png)
2、复制密文到bagua_de.go中shellcode变量
![image](https://user-images.githubusercontent.com/94209165/196104848-309c2271-db3b-489d-8731-e3b1849590b2.png)
3、执行bagua_de.go代码后成功上线
![image](https://user-images.githubusercontent.com/94209165/196105167-a423c576-97ea-40a8-a9fa-b13e943d7da3.png)
![image](https://user-images.githubusercontent.com/94209165/196105220-05868c4d-e88e-47cf-b06a-fb4e44c1f5fa.png)
4、编译后测试免杀效果。
![image](https://user-images.githubusercontent.com/94209165/196105259-bc83e505-868f-4558-95fe-fb0e9d52115a.png)
成功绕过火绒静态
![image](https://user-images.githubusercontent.com/94209165/196105274-af7a6ff3-ff64-4604-88dc-9b5d757c38c3.png)
成功绕过火绒动态
![image](https://user-images.githubusercontent.com/94209165/196105323-3f11ca8d-89f6-4da6-8760-90f0e57ef004.png)
---
---

主要分享一下免杀思路，目前只测试360和火绒动静态全过，其他自测。Golang小白，大佬们轻喷~
