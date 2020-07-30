<template>
  <div id="app">
    here is app div.
    <br>这里是路径
    <div id="exp">
        <el-dropdown split-button type="info" v-if="os == 'windows'" @command="handleCommand" @click="clickDriver">
            {{current_driver}}
            <el-dropdown-menu slot="dropdown">
                <el-dropdown-item v-for="i in driver" v-bind:command="i">{{ i }}</el-dropdown-item>
            </el-dropdown-menu>
        </el-dropdown>
        <el-dropdown split-button type="info" v-if="os != 'windows'" @command="handleCommand" @click="clickDriver">
            /
        </el-dropdown>
        <el-button-group>
            <el-button type="primary"  v-for="(i,index) in path" v-bind:val="index" @click="clickH($event)">{{ i }}<i class="el-icon-arrow-right el-icon--right"></i></el-button>
        </el-button-group>
        <table>
            <th>名称</th>
            <th>最后修改时间</th>
            <th>类型</th>
        <tr v-for="item in data" :key="item.fileName" v-bind:path="item.fileName" v-bind:typ="item.IsDir?1:2" @click="clickTrHandler($event)">
            <td ><img src="./assets/file_ico.png" width="16" height="16" v-if="!item.IsDir">
                <img src="./assets/dictory_ico.png" width="16" height="16"  v-if="item.IsDir">
                {{ item.fileName }}
            </td>
            <td>{{item.lastModTime.substring(0,19)}}</td>
            <td v-if="item.IsDir">文件夹</td>
            <td v-if="!item.IsDir">文件</td>
        </tr>
        </table>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import Vue from 'vue'
export default {
  name: 'App',
  created() {
    var that = this;
    console.log("in created func")
      axios.get("http://localhost:8000/getos")
          .then(function (response) {
              // console.log(response.data)
              that.os = response.data.os
              that.driver = response.data.data
              if(that.os != "windows"){
                  that.current_driver = [];
                  axios.get("http://localhost:8000/index/")
                      .then(function (response) {
                          // console.log(response.data.data)
                          that.data = response.data.data
                      })

              }else{
                  axios.get("http://localhost:8000/index/c")
                      .then(function (response) {
                          // console.log(response.data.data)
                          that.data = response.data.data
                      })
              }
          })


  },
  methods:{
      clickDriver(e){
          this.handleCommand(this.current_driver)
      },
      clickH(e){
        console.log("clickH")
        let that = this;
        let pos = e.target.getAttribute("val")||e.target.parentNode.getAttribute("val");
        this.path = this.path.slice(0,Number(pos)+1)
          let str = ""
        if(that.os == "windows"){
            str = "http://localhost:8000/index/"+this.current_driver.substring(0,1)+"/"+this.path.join("/")
        }else{
            str = "http://localhost:8000/index/"+this.path.join("/")
        }
          axios.get(str)
              .then(function (response) {
                  that.data =  response.data.data;
              })
      },
      clickTrHandler(e){
          let that = this;
          let path = e.target.parentNode.getAttribute("path")
          let typ = e.target.parentNode.getAttribute("typ")
          console.log(path)
          console.log(typ)
          console.log(e.target)
          // console.log(path)  // 路径
          // console.log(typ)              // 类型 1为文件夹，2为文件
          this.path.push(path)
          if(typ == "1"){
              let str = ""
              if (that.os == "windows"){
                  str = "http://localhost:8000/index/"+this.current_driver.substring(0,1)+"/"+this.path.join("/")
              }else{
                  str = "http://localhost:8000/index/"+this.path.join("/")
              }
              axios.get(str)
                  .then(function (response) {
                      that.data =  response.data.data;
                  })
          }else{
              let str = ""
              if(that.os == "windows"){
                  str = "http://localhost:8000/file/"+this.current_driver.substring(0,1)+"/"+this.path.join("/")
              }else{
                  str = "http://localhost:8000/file/"+this.path.join("/")
              }
              window.open(str)
              this.path.pop()
          }

      },
      handleCommand(command) {
          let that = this;
          this.$message('click on item ' + command);
          let str = ""
          if(that.os == "windows"){
              this.current_driver = command;
              str = "http://localhost:8000/index/"+this.current_driver.substring(0,1)
          }else{
              str = "http://localhost:8000/index/"
          }
              this.path = [];
              axios.get(str)
                  .then(function (response) {
                      that.data =  response.data.data;
                  })
      }
  },
  data(){
    return {
      msg:"hha",
        os:"",
      data:"",
      current_driver:'C:',
      driver:"",
      path:[]
    }
  },
  components: {
  }
}
</script>

<style>
#exp {
  border: solid black 1px;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

ul {
    list-style: none;
}
tr:hover td{
    background-color: #eee;
}
/*.el-dropdown-link {*/
/*    cursor: pointer;*/
/*    color: #409EFF;*/
/*}*/
/*.el-icon-arrow-down {*/
/*    font-size: 12px;*/
/*}*/
</style>
