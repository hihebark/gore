<p align="center"><img src="Cover1.jpg" width="500"></p>

gore
=====
gore is a tool to detect (and recognize) face (shape?) in an image, without the use of OpenCV or another library. it use HOG [Histogram of Oriented Gradients](https://en.wikipedia.org/wiki/Histogram_of_oriented_gradients), further link to check:
* [Machine Learning is Fun! Part 4: Modern Face Recognition with Deep Learning](https://medium.com/@ageitgey/machine-learning-is-fun-part-4-modern-face-recognition-with-deep-learning-c3cffc121d78)
* [HOG Person Detector Tutorial](http://mccormickml.com/2013/05/09/hog-person-detector-tutorial/)
* [Gradient Vectors](http://mccormickml.com/2013/05/07/gradient-vectors/)

Result
-------
#### Almost Done
Not done yet but hope to get into this result.
<p align="center"><img src="GoreProject.png" width="400"></p>

#### Result so far!
using HOG implementation.
<p align="center"><img src="face-hog.png" width="400"></p>

#### Usage

`go run main.go -img data/image.png`


Steps
------

- [x] Grayscal the image.
- [x] Draw square.
- [x] Divid the image into 16*16 cells.
- [x] HOG implementation.
- [ ] Find face.
- [x] Output the image surrounding the face(s) with rectangle box. ~~(almost)~~
- [ ] Analyze the face.
- [ ] Compare.
- [ ] Predection.

