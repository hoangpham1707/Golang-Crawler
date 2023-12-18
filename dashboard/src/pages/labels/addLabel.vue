<!-- ListSource.vue -->
<template>
    <div class="container mx-auto p-4">
        <h1 class="text-2xl font-semibold mb-4">List mention</h1>
        <div>
            Tổng Web chưa gán: <p>{{ totalWeb.total }}</p>
            Tổng Facebook chưa gán: <p>{{ totalFace.total }}</p>
        </div>
        <div>
            Tiêu cực: <p>{{ totalNeg }}</p>
            Tích cực: <p>{{ totalPos }}</p>
            Trung tính: <p>{{ totalNeu }}</p>
        </div>


    </div>
</template>

  
<script>
import axios from 'axios';

export default {
    name: 'ListAddLabel',
    data() {
        return {
            totalNeg: 0,
            totalNeu: 0,
            totalPos: 0,
            totalFace: 0,
            totalWeb: 0,
        };
    },
    created() {
        this.fetchMentions()
        this.fetchMentionsCount()
    },
    methods: {

        async fetchMentionsCount() {
            try {
                const response = await axios.request({
                    method: 'GET',
                    url: `http://localhost:6006/api/v1/common/mentions/count`
                });
                console.log({ response });
                this.totalFace = response.data.data.totalFacebook;
                this.totalWeb = response.data.data.totalWeb;

                //  console.log('Current Page:', this.totalWeb);
            } catch (error) {
                console.error('Error fetching sources:', error);
            }
        },
        async fetchMentions() {
            try {
                const response = await axios.request({
                    method: 'GET',
                    url: `http://localhost:6006/api/v1/common/mentions/countLabel`
                });
                console.log({ response });
                this.totalNeg = response.data.data.Neg;
                this.totalNeu = response.data.data.Neu;
                this.totalPos = response.data.data.Pos;
                //  console.log('Current Page:', this.totalWeb);
            } catch (error) {
                console.error('Error fetching sources:', error);
            }
        }
    },
};
</script>
  
<style scoped>
/* Add styling for the entry container */
.entry-container {
    max-height: 110px;
    /* Adjust the height as needed */
    overflow: hidden;
    border: 1px solid #ddd;
    padding: 10px;
    margin-bottom: 10px;
}

.entry-list {
    display: flex;
    flex-wrap: wrap;
    /* Optional: Wrap to the next line if the items don't fit */
    justify-content: space-between;
    /* Optional: Adjust the alignment */
    list-style: none;
    /* Optional: Remove default list styling */
    padding: 0;
    /* Optional: Remove default padding */
}

.pagination {
    display: flex;
    justify-content: space-between;
}
</style>
