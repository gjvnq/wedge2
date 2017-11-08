<template>
  <div class="row">
    <div class="col-md-12">
      <div class="card">
        <div class="header">
          <h4 class="title">{{$t('Add Currency or Asset')}}</h4>
        </div>
        <div class="content">
          <form lpformnum="1">
            <div class="row">
              <div class="col-md-3">
                <div class="form-group">
                  <label>{{$t('Id')}}</label>
                  <input type="text" class="form-control border-input">
                </div>
              </div>
              <div class="col-md-5">
                <div class="form-group">
                  <label>{{$t('Name')}}</label>
                  <input type="text" class="form-control border-input">
                </div>
              </div>
              <div class="col-md-2">
                <div class="form-group">
                  <label>{{$t('Decimal Places')}}</label>
                  <input type="number" class="form-control border-input">
                </div>
              </div>
              <div class="col-md-2">
                <div class="form-group">
                  <label>{{$t('Formating')}}</label>
                  <input type="text" class="form-control border-input">
                </div>
              </div>
            </div>
            <div class="text-center">
              <button type="submit" class="btn btn-info btn-fill btn-wd">{{$t('Add Currency or Asset')}}</button>
            </div>
            <div class="clearfix">
            </div>
          </form>
        </div>
      </div>
    </div>
    <div class="col-md-12">
      <div class="card">
        <paper-table :title="$t(table1.title)" :sub-title="$t(table1.subTitle)" :data="table1.data" :columns="table1.columns">

        </paper-table>
      </div>
    </div>
  </div>
</template>
<script>
  import PaperTable from 'components/UIComponents/PaperTable.vue'
  const tableColumns = ['Id', 'Name', 'Last Value', '']
  const tableData = [{
    id: 1,
    name: 'Dakota Rice',
    salary: '$36.738',
    country: 'Niger',
    city: 'Oud-Turnhout'
  },
  {
    id: 2,
    name: 'Minerva Hooper',
    salary: '$23,789',
    country: 'Curaçao',
    city: 'Sinaai-Waas'
  },
  {
    id: 3,
    name: 'Sage Rodriguez',
    salary: '$56,142',
    country: 'Netherlands',
    city: 'Baileux'
  },
  {
    id: 4,
    name: 'Philip Chaney',
    salary: '$38,735',
    country: 'Korea, South',
    city: 'Overland Park'
  },
  {
    id: 5,
    name: 'Doris Greene',
    salary: '$63,542',
    country: 'Malawi',
    city: 'Feldkirchen in Kärnten'
  }]

  export default {
    components: {
      PaperTable
    },
    data () {
      return {
        table1: {
          title: 'Currencies & Assets',
          subTitle: '',
          columns: [...tableColumns],
          data: [...tableData]
        }
      }
    }
  }

</script>
<style>

</style>
