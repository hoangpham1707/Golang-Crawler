<template>
    <img src='../../assets/images/ptit.png' alt='logo' height='40' />
    <span class='color-primary' style='font-weight:600' v-if='innerWidth > 991'>
        PTIT CMS
    </span>
    <span class='color-primary' v-else>
        PTIT
    </span>
</template>
  
<script>
export default {
    name: 'NavBar',
    data: () => ({
        innerWidth: 0
    }),
    methods: {},
    created() {
        if (process.client) {
            this.innerWidth = window.innerWidth
        }
    },
    mounted() {
    }
}
</script>
  
<style scoped></style>
  