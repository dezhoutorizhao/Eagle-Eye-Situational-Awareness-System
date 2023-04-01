import argparse
import os
import platform
import base64
import sys
import json
from pathlib import Path
from datetime import datetime

import torch

FILE = Path(__file__).resolve()
ROOT = FILE.parents[0]  # YOLOv5 root directory
if str(ROOT) not in sys.path:
    sys.path.append(str(ROOT))  # add ROOT to PATH
ROOT = Path(os.path.relpath(ROOT, Path.cwd()))  # relative

from models.common import DetectMultiBackend
from utils.dataloaders import IMG_FORMATS, VID_FORMATS, LoadImages, LoadScreenshots, LoadStreams
from utils.general import (LOGGER, Profile, check_file, check_img_size, check_imshow, check_requirements, colorstr, cv2,
                           increment_path, non_max_suppression, print_args, scale_boxes, strip_optimizer, xyxy2xywh)
from utils.plots import Annotator, colors, save_one_box
from utils.torch_utils import select_device, smart_inference_mode
import numpy as np

class Detect():
    def __init__(self, src, conf, interval, location, task, webcame=True):
        self.weights = 'Fire.pt'
        self.img_size = 640
        self.conf_thres =float(conf)
        self.iou_thres =0.5
        self.detect_interval = int(interval)
        self.currentFrame = 0
        self.location = location
        self.task = task

        if torch.cuda.is_available():
            self.device = torch.device("cuda:0")
        else:
            self.device = torch.device("cpu")
        self.half = self.device.type == 'cuda' and torch.cuda.is_available()

        self.model = DetectMultiBackend(self.weights, device=self.device)

        self.model.to(self.device).eval()
        self.model.float()  # to FP16
        self.names = self.model.module.names if hasattr(self.model, 'module') else self.model.names
        # Run inference
        img_init = torch.zeros((1, 3, self.img_size, self.img_size), device=self.device)  # init img
        _ = self.model(img_init.half() if self.half else img_init) if self.device.type != 'cpu' else None  # run once
        self.webcamera = webcame
        if self.webcamera:
            self.dataset = LoadStreams(src, img_size=self.img_size, stride=self.model.stride, auto=self.model.pt, vid_stride=1)
        else:
            self.dataset = LoadImages(src)

    def detect(self):

        detect_counter = 0
        for path, im, im0s, vid_cap, s in self.dataset:
            self.currentFrame += 1

            # 检查检测计数器是否已达到检测间隔
            if detect_counter < self.detect_interval:
                detect_counter += 1
                continue
            else:
                detect_counter = 0

            im = torch.from_numpy(im).to(self.model.device)
            im = im.float()  # if self.model.fp16 else im.float()  # uint8 to fp16/32
            im /= 255  # 0 - 255 to 0.0 - 1.0
            if len(im.shape) == 3:
                im = im[None]  # expand for batch dim
            pred = self.model(im, augment=False, visualize=False)
            pred = non_max_suppression(pred, self.conf_thres, self.iou_thres)
            resconf = []
            for i, det in enumerate(pred):
                if self.webcamera:
                    p, im0, frame = path[i], im0s[i].copy(), self.dataset.count
                    s += f'{i}: '
                else:
                    p, im0, frame = path, im0s.copy(), getattr(self.dataset, 'frame', 0)
                annotator = Annotator(im0, line_width=3, example=str(self.names))
                if len(det):
                    det[:, :4] = scale_boxes(im.shape[2:], det[:, :4], im0.shape).round()
                    for *xyxy, conf, cls in reversed(det):
                        c = int(cls)  # integer class
                        label = self.names[c].capitalize()
                        annotator.box_label(xyxy, label, color=colors(c, True))
                        resconf.append(float(conf))
                img0 = annotator.result()

                if self.currentFrame > 0 and any(c > self.conf_thres for c in resconf):                     
                    if not os.path.exists('detect'):                             
                        os.makedirs('detect')                     
                    time_now = datetime.now().strftime("%Y-%m-%d_%H-%M-%S-%f")                          
                    img_path = f"./detect/{time_now}.jpg"                          
                    cv2.imwrite(img_path, img0)                    
                    # 获取识别结果                     
                    rate = sum(resconf) / len(resconf) if len(resconf) > 0 else 0                     
                    detected = any(c > self.conf_thres for c in resconf)                     
                    result = {                         
                        "Photo": img_path,                         
                        "Rate": str(rate),                         
                        "Task": self.task,                         
                        "Location": self.location,                     
                    }                      
                    if detected:                                  
                        json_str = json.dumps(result)                          
                        print(json_str)                                     
                        return json_str                    
                    else:                         
                        continue
