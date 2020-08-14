#!/bin/zsh

if [ -e ./testdata ];then
chmod -R 755 ./testdata
fi
rm -r testdata
mkdir testdata

# 正常系
mkdir testdata/img
mkdir testdata/img/img

curl http://flat-icon-design.com/f/f_object_174/s512_f_object_174_0bg.jpg > ./testdata/img/azarashi.jpg
curl http://flat-icon-design.com/f/f_object_174/s512_f_object_174_0bg.jpg > ./testdata/img/img/azarashi.jpg
curl http://flat-icon-design.com/f/f_object_149/s512_f_object_149_0bg.jpg > ./testdata/img/tanuki.jpg
curl http://flat-icon-design.com/f/f_object_149/s512_f_object_149_0bg.jpg > ./testdata/img/img/tanuki.jpg
curl http://flat-icon-design.com/f/f_object_157/s512_f_object_157_0bg.png > ./testdata/img/osaru.png


# 異常系
mkdir testdata/err

mkdir testdata/err/read_permission
curl http://flat-icon-design.com/f/f_object_149/s512_f_object_149_0bg.jpg > ./testdata/err/read_permission.jpg
chmod 355 ./testdata/err/read_permission
chmod 355 ./testdata/err/read_permission.jpg

mkdir testdata/err/write_permission
curl http://flat-icon-design.com/f/f_object_149/s512_f_object_149_0bg.jpg > ./testdata/err/write_permission/write_permission.jpg
chmod 555 ./testdata/err/write_permission
chmod 555 ./testdata/err/write_permission/write_permission.jpg
