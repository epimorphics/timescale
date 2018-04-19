<template>
  <div>
    <el-table
      :data="tableData"
      style="width: 100%">
      <el-table-column
        prop="name"
        label="Project"
        >
        <template slot-scope="scope">
          <el-input v-model="scope.row.name" v-if="scope.row.editing">
          </el-input>
          <div v-else>
            {{ scope.row.name }}
          </div>
        </template>
      </el-table-column>
      <el-table-column
        prop="person"
        label="Person"
        >
        <template slot-scope="scope">
          <el-input v-model="scope.row.person" v-if="scope.row.editing">

          </el-input>
          <div v-else>
            {{ scope.row.person}}
          </div>
        </template>
      </el-table-column>
      <el-table-column
        prop="weeks"
        label="Weeks"
        >
        <template slot-scope="scope">
          <div v-if="scope.row.editing">
            <el-input-number v-model="scope.row.weeks">
            </el-input-number>
          </div>
          <div v-else>
            {{ scope.row.weeks }}
          </div>
        </template>
      </el-table-column>
      <el-table-column>
        <template slot-scope="scope">
          <div v-if="scope.row.editing">
            <el-button @click="save(scope.row)">Save</el-button>
            <el-button @click="cancel(scope.row.id)">Cancel</el-button>
            <el-button @click="deleteProject(scope.row)">Delete</el-button>
          </div>
          <el-button v-else @click="editing(scope.row.id)">Edit</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-button @click="create">Add New Project</el-button>
  </div>
</template>
<script>
import request from 'superagent'
import Cookies from 'js-cookie'
export default {
  name: 'tasktable',
  mounted () {
    this.getTable()
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
    },
    editing(id) {
      var vueM = this
      this.tableData = this.tableData.map((row) => {
        if (row.id == id) {
          row.editing = true
          vueM.savedRow = JSON.parse(JSON.stringify(row))
        } 
        return row
      })
    },
    deleteProject(row) {
      var vueM = this
      if (this.isNew(row)) {
        this.tableData = this.tableData.filter((row) => !vueM.isNew(row))
        return
      }
      request.del(`/projects/${row.id}`)
        .set({'Content-Type': 'application/json', 'Authorization': 'Bearer ' +  Cookies.get('token')})
        .then(() => {
          vueM.getTable()
          vueM.$emit('update') 
        })
        .catch(() => {
          console.log('failed')
        })
    },
    save(row) {
      var vueM = this
      var saveRequest
      if (this.isNew(row)) {
        saveRequest = request.post('/projects')
      } else {
        saveRequest = request.put('/projects')
      }
      saveRequest
        .set({'Content-Type': 'application/json', 'Authorization': 'Bearer ' + Cookies.get('token')})
        .send(row)
        .then(() => {
          vueM.getTable()
          vueM.$emit('update') 
        })
        .catch((err) => {
          console.log(err)
          this.cancel(row.id)
        })
    },
    create() {
      this.tableData.push({
        name: '',
        person: '',
        days: 5,
        editing: true
      })
    },
    isNew(row) {
      return !row.hasOwnProperty('id')
    },
    cancel(id) {
      var vueM = this
      this.tableData = this.tableData.filter((row) => !vueM.isNew(row))
      this.tableData = this.tableData.map((row) => {
        if (id == row.id) {
          row = vueM.savedRow
        }
        row.editing = false
        return row
      })
    }
  },
  data () {
    return {
      tableData: [],
      savedRow: {},
    }
  }
}
</script>
<style>
</style>
