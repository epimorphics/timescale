<template>
  <div class="task-container">
    <taskcalendar :tasks="tableData" class="calendar"/>
    <tasktable @update="getTable" class="table"/>
  </div>
</template>
<script>
import taskcalendar from '@/components/TaskCalendar'
import tasktable from '@/components/TaskTable'
import Cookies from 'js-cookie'
import request from 'superagent'
export default {
  name: 'timescale',
  components: {
    tasktable,
    taskcalendar
  },
  data () {
    return {
      tableData: []
    }
  },
  methods: {
    getTable() {
      var vueM = this
      request.get('/projects')
        .set({'Content-Type': 'application/json', 'Authorization': 'Bearer ' + Cookies.get('token')})
        .then((resp) => {
          if (resp.body == null) {
            vueM.tableData = []
            return
          }
          vueM.tableData = resp.body
        })
    }
  },
  mounted () {
    if (!Cookies.get('token')) {
      this.$router.push({ path: 'login' })
    } 
    this.getTable()
  }
}
</script>
<style>
.calendar {
  max-height: 20em;
}

.task-container {
  width: 85vw;
  margin: 2em auto 0 auto;
}

.table {

  
}
</style>
