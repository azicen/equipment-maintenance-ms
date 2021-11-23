<template>
  <div>
    <h2>
      设备维护管理系统
      <div class="word-rotation-text">
        <span data-show>Starting on 2021.11.22</span>
        <span>BY strluck AND null122</span>
        <span>Completed in 2021.xx.xx</span>
        <span>xxxxxxxxxxxxxxxxxxxxxxx</span>
      </div>
    </h2>
  </div>
</template>

<script>
export default {
  name: 'WordRotation',
  components: {},
  created() {
    setInterval(this.RotationTime, 3000);
  },
  beforeCreate() {
    clearInterval(this.RotationTime);
  },
  methods: {
    RotationTime() {
      const show = document.querySelector('span[data-show]');
      const next = show.nextElementSibling || document.querySelector('span:first-child');
      const up = document.querySelector('span[data-up]');

      if (up) {
        up.removeAttribute('data-up');
      }

      show.removeAttribute('data-show');
      show.setAttribute('data-up', '');

      next.setAttribute('data-show', '');
    }
  }
}
</script>

<style>

body {
  display: flex;
  justify-content: center;
  /*align-items: center;*/
  min-height: 100vh;
}

h2 {
  width: 1024px;
  font-size: 88px;
  font-family: Helvetica;
  /*line-height: 1.06;*/
  letter-spacing: -0.02em;
  color: #1d1d1f;
}

.word-rotation-text {
  height: 106px;
  position: relative;
  overflow: hidden;
}

.word-rotation-text span {
  display: block;
  box-sizing: border-box;
  position: absolute;
  top: 100px;

  background-size: 100% 100%;
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-repeat: no-repeat;
}

.word-rotation-text span[data-show] {
  transform: translateY(-100%);
  transition: .5s transform ease-in-out;
}

.word-rotation-text span[data-up] {
  transform: translateY(-200%);
  transition: .5s transform ease-in-out;
}

.word-rotation-text span:nth-child(1) {
  background-image: linear-gradient(45deg, #0ecffe 50%, #07a6f1);
}

.word-rotation-text span:nth-child(2) {
  background-image: linear-gradient(45deg, #18e198 50%, #0ec15d);
}

.word-rotation-text span:nth-child(3) {
  background-image: linear-gradient(45deg, #8a7cfb 50%, #633e9c);
}

.word-rotation-text span:nth-child(4) {
  background-image: linear-gradient(45deg, #fa7671 50%, #f45f7f);
}
</style>