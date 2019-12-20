model-generator
===

this is a model generator for [jinzhu/gorm](https://github.com/jinzhu/gorm)

Demo
===
![image](https://github.com/bigkucha/model-generator/blob/master/media/test.gif)


Install
===

```
$ go get github.com/bigkucha/model-generator
```

Usage
===

```
$ $GOPATH/bin/model-generator -u=root -p=(pwd of your mysql) -d=database -t=table -dir=
```

Flags
===
<table>
 <tr>
    <th>Flag</th>
    <th>Rule</th>
    <th>Usage</th>
  </tr>
  <tr>
    <td>username, u</td>
    <td>optional, default 'root'</td>
    <td>username of mysql</td>
  </tr>
  
  <tr>
    <td>password, p</td>
    <td>require, default null</td>
    <td>password of mysql</td>
  </tr>
  
  <tr>
    <td>database, d</td>
    <td>require</td>
    <td>select your database</td>
  </tr>
  <tr>
    <td>table, t</td>
    <td>optional,default 'ALL'</td>
    <td>chose table to generate model, if not set ,all tables in your database will be used</td>
  </tr>
  
  <tr>
    <td>dir</td>
    <td>optional,default 'model' of current directory</td>
    <td>models where to be store</td>
  </tr>
</table>

Note
===
 run `$GOPATH/bin/model-generator -h` for more help

TODO
===
- Multi tables
- regex 