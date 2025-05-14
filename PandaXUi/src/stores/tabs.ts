import { defineStore } from "pinia";

export const useTabsStore = defineStore('tabs', {
    state:()=>({
        tabs:[], //存储所有打开的标签页
        activaTab:'' //当前激活的标签页
    }),
    actions:{
        //添加标签页
        addTab(route: { path: any; meta: any; name: any; }) {
            const { path, meta, name } = route
            const tabExists = this.tabs.some(tab => tab.path === path)
            
            if (!tabExists) {
              this.tabs.push({
                title: meta.title || name,
                path,
                name,
                closable: true
              })
            }
            this.activaTab = path
          },
          
          // 关闭标签页
          closeTab(path: any) {
            const index = this.tabs.findIndex(tab => tab.path === path)
            if (index >= 0) {
              this.tabs.splice(index, 1)
            }
          },
          
          // 设置当前激活的标签页
          setActiveTab(path: any) {
            this.activaTab = path
          }
    }
})