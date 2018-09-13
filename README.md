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

##### Output

```

  ┏ ┳ ┓
  ┣ o ┫
  ┗ ┻ ┛
  Gore - 0.0.1
Usage of ./gore:
  -p string
        Path to the image.
  -s int
        Scale image into the given s. (default 2)

```

---

<p align="center"><i>Made with </i>■ <i>by <b>hihebark</b></i></p>
