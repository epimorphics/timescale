<template>
  <div class="border">
    <el-row>
      <el-col :span="6">
        month
      </el-col>
      <el-col class="border-left"  :span="6" v-for="month in months" :key="month.name">
        <el-row>{{ month.name }}</el-row>
        <el-row class="border-bottom" type="flex" justify="space-around">
          <el-col  v-for="monday in month.mondays" :key="`${month} ${monday}`" :span="2">{{ monday }}</el-col>
        </el-row>
      </el-col>
    </el-row>
    <el-row class="border-bottom" v-for="person in people" :key="person">
      <el-col :span="6">
        <el-row class="task">{{ person }}</el-row>
      </el-col>
      <el-col v-for="task in personTasks(person)" :key="task.id" :span="task.weeks">
        <div class="task" :style="`background: ${task.color};`">{{ task.name }}</div>
      </el-col>
    </el-row>
  </div>
</template>
<script>
import moment from 'moment'
export default {
  name: 'taskcalendar',
  props: ['tasks'],
  data () {
    return {
      start: moment().weekday(0).hour(1).date(),
      people: [],
      months: ['april', 'may', 'june'],
      colors: ['rgb(128,256,128)','rgb(128,128,256)','rgb(256,128,128)']
    }
  },
  mounted () {
    this.getPeople()
    this.getMonths()
  },
  watch: {
    tasks () {
      this.getPeople()
    }
  },
  methods: {
    personTasks(name) {
      var vueM = this
      var newtasks = this.tasks.filter((task) => task.person == name )
      return newtasks.map((task) => {
        task.color = vueM.colors[task.id % 3]
        return task
      })
    },
    getPeople () {
      this.people = this.tasks.reduce((acc, task) => { 
        if (!acc.includes(task.person)) {
          acc.push(task.person) 
        }
        return acc
      }, [])
    },
    getMonths () {
      var months = []
      for (var i = 0; i < 3; i++) {
        var monthObj = {
          name: "",
          mondays: []
        }
        monthObj.name = moment().add(i, "month").format("MMMM")
        var monday = moment().add(i, "month").startOf('month').day('Monday')
        if (monday.date() > 7) { monday.add(7, 'd') }
        var month = monday.month()
        while(monday.month() === month) {
          monthObj.mondays.push(monday.date())
          monday.add(7, 'd')
        }
        months.push(monthObj)
      }
      this.months = months
    }
  }
}
</script>
<style>
  .task {
    width: 100%;
    padding: 2em;
  }
  .border {
    border: 1px solid rgb(64,64,64) ;
  }
  .border-bottom {
    border-top: 1px solid rgb(64,64,64) ;
  }
  .border-left {
    border-left: 1px solid rgb(64,64,64) ;
  }
</style>
