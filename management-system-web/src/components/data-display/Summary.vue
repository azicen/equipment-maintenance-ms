<template>
  <div class="summary-box-card">
    <el-card class="summary-box-card">
      <template #header>
        <div>
          <span>概览</span>
        </div>
      </template>
      <div>
        <div class="summary-item">
          <el-button class="summary-button">
            <p class="summary-item-label">今日维护</p>
            <p class="summary-item-value" id="today">{{ this.data.today }}</p>
          </el-button>
        </div>
        <div class="summary-item">
          <el-button class="summary-button">
            <p class="summary-item-label">等待维护</p>
            <p class="summary-item-value" id="wait">{{ this.data.wait }}</p>
          </el-button>
        </div>
        <div class="summary-item">
          <el-button class="summary-button">
            <p class="summary-item-label">报警设备</p>
            <p class="summary-item-value" id="alarm">{{ this.data.alarm }}</p>
          </el-button>
        </div>
        <div class="summary-item">
          <el-button class="summary-button">
            <p class="summary-item-label">临期报废</p>
            <p class="summary-item-value" id="scrap">{{ this.data.scrap }}</p>
          </el-button>
        </div>
      </div>
    </el-card>
  </div>
</template>

<script lang="ts">
import {defineComponent, PropType} from 'vue'

export interface SummaryData {
  today: number
  wait: number
  alarm: number
  scrap: number
}

export default defineComponent({
  name: 'Summary',
  props: {
    data: {
      type: Object as PropType<SummaryData>,
      required: true,
    }
  },
  setup() {
  }
})
</script>

<style>
.summary-box-card {
  /*width: 100%;*/
}

.summary-item {
  margin: 10px;
  display: inline;
}

.summary-button {
  height: 128px;
  width: 21%;
  max-width: 180px
}

.summary-item-label {
  font-size: 18px;
  margin: 8%;
}

.summary-item-value {
  font-size: 30px;
}

#today {
  color: yellowgreen;
  font-weight: normal;
}

#wait {
  color: dimgrey;
  font-weight: normal;
}

#alarm {
  color: salmon;
  font-weight: normal;
}

#scrap {
  color: crimson;
  font-weight: 900;
}
</style>