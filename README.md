# geoip
[![Build Status](https://travis-ci.org/gonet2/geoip.svg?branch=master)](https://travis-ci.org/gonet2/geoip)

# 设计思路
查询IP归属地，基于maxmind的geoip2库做的封装，如果需要最新的准确的数据，需要向maxmind购买。     
(query geo-locations of IP, if you need accurate & updated data, please purchase from maxmind.com, thanks. )        

> 问: 为什么选择maxmind的geoip2库？         
> 答: maxmind的geoip2的库设计为一个支持mmap的二叉树文件，查询时间复杂度为O(logN),        文件大小不超过100M，极其紧凑，省内存，速度快，零配置，是目前见过的最好的方案。                

## 使用
参考测试用例

## 安装
参考Dockerfile
