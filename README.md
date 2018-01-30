[![Build Status](https://travis-ci.org/kyle-wang/SLGFov.svg?branch=master)](https://travis-ci.org/kyle-wang/SLGFov)

说明
*  这是一个为SLG游戏的大地图可视范围内的行军行为提供基础计算函数的golang库，默认所有的坐标点都处于第一象限

Package说明

* CheckIntersectSquareAndBeeline 函数实现判断一条行军线是否穿过某个视野区域
    *  Chart结构体定义了起点和终点（对角线点），直线和矩形都用同样的结构表示
    *  如果直线穿过定义的矩形区域，返回true，否则返回false
* Bresenham 函数计算一条直线经过哪些坐标块
    *  函数返回值[]PointCoordinate，返回经过的所有坐标块
    *  输入参数不限定直线方向，但是依然只能是处于第一象限的坐标

使用方法
*  go get github.com/kyle-wang/SLGFov
*  SLGFov.CheckIntersectSquareAndBeeline(squre,beeline)
*  SLGFov.Bresenham(x1, y1, x2, y2)
    *  x1,y1为起点坐标，x2,y2为终点坐标