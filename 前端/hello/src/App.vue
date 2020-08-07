<template>
  <div id="app">
    here is app div.
    <br />这里是路径
    <div id="exp">
      <el-dropdown
        split-button
        type="info"
        v-if="os == 'windows'"
        @command="handleCommand"
        @click="clickDriver"
      >
        {{current_driver}}
        <el-dropdown-menu slot="dropdown">
          <el-dropdown-item v-for="i in driver" v-bind:command="i" :key="i">{{ i }}</el-dropdown-item>
        </el-dropdown-menu>
      </el-dropdown>
      <el-dropdown
        split-button
        type="info"
        v-if="os != 'windows'"
        @command="handleCommand"
        @click="clickDriver"
      >/</el-dropdown>
      <el-button-group>
        <el-button
          type="primary"
          v-for="(i,index) in path"
          v-bind:val="index"
          @click="clickH($event)"
          :key="index"
        >
          {{ i }}
          <i class="el-icon-arrow-right el-icon--right"></i>
        </el-button>
      </el-button-group>
      <table>
        <th>名称</th>
        <th>最后修改时间</th>
        <th>类型</th>
        <tr
          v-for="item in data"
          :key="item.fileName"
          v-bind:path="item.fileName"
          v-bind:typ="item.IsDir?1:2"
          @click="clickTrHandler($event)"
          @mouseover="m_over($event)"
        >
          <td>
            <img src="./assets/file_ico.png" width="16" height="16" v-if="!item.IsDir" />
            <img src="./assets/dictory_ico.png" width="16" height="16" v-if="item.IsDir" />
            {{ item.fileName }}
          </td>
          <td>{{item.lastModTime.substring(0,19)}}</td>
          <td v-if="item.IsDir">文件夹</td>
          <td v-if="!item.IsDir">文件</td>
        </tr>
      </table>
    </div>
    <!--  右键菜单部分 -->
    <div v-show="con" id="pos">
      <ul class="menu">
        <li @click="m_open">打开</li>
        <li>剪切</li>
        <li>复制</li>
        <li>删除</li>
        <li>重命名</li>
        <li>属性</li>
      </ul>
    </div>
  </div>
</template>

<script>
import Vue from "vue";
import interactive from "./utils/interactive";
export default {
  name: "App",
  created() {
    var that = this;
    console.log("in created func");
    //此处URL不同平台改变
    interactive.init("/getos", that);


    // 右键禁止以及自定义菜单弹出
    document.oncontextmenu = function (e) {
      that.con = true;
      let dom_menu = document.getElementById("pos");
      // console.log(dom_menu);
      dom_menu.style.left = e.clientX + "px";
      dom_menu.style.top = e.clientY - 10 + "px";
      return false;
    };
  },
  methods: {
    m_over(e){
      let tat = e.target.parentNode;
      let path = tat.getAttribute("path");
      let typ = tat.getAttribute("typ");
      this.m_target.path = path;
      this.m_target.typ = typ;
    },
    m_open() {
      this.con = false;
      let path = this.m_target.path;
      let typ = this.m_target.typ;
      var that = this;
       this.path.push(path);
      if (typ == "1") {
        let str = "";
        if (that.os == "windows") {
          str =
            "/index/" +
            this.current_driver.substring(0, 1) +
            "/" +
            this.path.join("/");
        } else {
          str = "/index/" + this.path.join("/");
        }
        interactive.getData(str, that);
      } else {
        let str = "";
        if (that.os == "windows") {
          str =
            interactive.baseURL +
            "/file/" +
            this.current_driver.substring(0, 1) +
            "/" +
            this.path.join("/");
        } else {
          str = interactive.baseURL + "/file/" + this.path.join("/");
        }
        window.open(str);
        this.path.pop();
      }
    },
    clickDriver(e) {
      this.handleCommand(this.current_driver);
    },
    clickH(e) {
      console.log("clickH");
      let that = this;
      let pos =
        e.target.getAttribute("val") || e.target.parentNode.getAttribute("val");
      this.path = this.path.slice(0, Number(pos) + 1);
      let str = "";
      if (that.os == "windows") {
        str =
          "/index/" +
          this.current_driver.substring(0, 1) +
          "/" +
          this.path.join("/");
      } else {
        str = "/index/" + this.path.join("/");
      }
      interactive.getData(str, that);
    },
    clickTrHandler(e) {
      let that = this;
      let path = e.target.parentNode.getAttribute("path");
      let typ = e.target.parentNode.getAttribute("typ");
      console.log(path);
      console.log(typ);
      console.log(e.target);
      // console.log(path)  // 路径
      // console.log(typ)              // 类型 1为文件夹，2为文件
      this.path.push(path);
      if (typ == "1") {
        let str = "";
        if (that.os == "windows") {
          str =
            "/index/" +
            this.current_driver.substring(0, 1) +
            "/" +
            this.path.join("/");
        } else {
          str = "/index/" + this.path.join("/");
        }
        interactive.getData(str, that);
      } else {
        let str = "";
        if (that.os == "windows") {
          str =
            interactive.baseURL +
            "/file/" +
            this.current_driver.substring(0, 1) +
            "/" +
            this.path.join("/");
        } else {
          str = interactive.baseURL + "/file/" + this.path.join("/");
        }
        window.open(str);
        this.path.pop();
      }
    },
    handleCommand(command) {
      let that = this;
      this.$message("click on item " + command);
      let str = "";
      if (that.os == "windows") {
        this.current_driver = command;
        str = "/index/" + this.current_driver.substring(0, 1);
      } else {
        str = "/index/";
      }
      this.path = [];
      interactive.getData(str, that);
    },
  },
  data() {
    return {
      con: false,
      msg: "hha",
      os: "",
      data: "",
      current_driver: "C:",
      driver: "",
      path: [],
      m_target: {
        path:"",
        typ:"",
      },
    };
  },
  components: {},
};
</script>

<style>
* {
  margin: 0;
  padding: 0;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
table {
  width: 100%;
}
/* -------右键菜单 */
#exp {
  border: solid black 1px;
}
ul {
  list-style: none;
}
tr:hover td {
  background-color: #eee;
}

.menu {
  height: 130px;
  width: 80px;
  /* border-radius: 10px; */
  border: 1px solid #999999;
  background-color: #f4f4f4;
}
.menu li {
  display: block;
  line-height: 20px;
}
.menu li:hover {
  background-color: #1790ff;
  color: white;
}
#pos {
  position: fixed;
  top: 50px;
  left: 10px;
}
</style>
