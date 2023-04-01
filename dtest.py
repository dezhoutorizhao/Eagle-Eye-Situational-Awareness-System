# -*- coding: utf-8 -*-
from detect_cls import Detect
import cv2
import base64
import json

import argparse      

parser = argparse.ArgumentParser(description='YOLOv5 Object Detection') 
parser.add_argument('--src', type=str, default='0', help='source') # 视频或图像的路径或摄像头的索引   
parser.add_argument('--conf', type=float, default=0.25, help='object confidence threshold')   # 目标置信度阈值  
parser.add_argument('--interval', type=int, default=10, help='detection interval')  # 目标检测的间隔帧数  
parser.add_argument('--location', type=str, default='Unknown', help='location of detection') # 检测场景的位置  
parser.add_argument('--task', type=str, default='Detection', help='detection task name')  # 检测任务的名称  
parser.add_argument('--webcam', action='store_true', help='use webcam as source')  # 是否使用摄像头作为源   

args = parser.parse_args()    

dd = Detect(args.src, args.conf, args.interval, args.location, args.task, args.webcam)
xx = dd.detect()

