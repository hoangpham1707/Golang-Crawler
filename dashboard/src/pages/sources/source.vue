<template>
    <div class="container mx-auto p-1">
        <p class="text-2xl font-semibold mb-1">Danh sách nguồn</p>
        <table class="w-full border">
            <thead>
                <tr class="bg-gray-200">
                    <th class="border px-4 py-2">Loại nguồn</th>
                    <th class="border px-4 py-2">Số lượng bài đề cập</th>
                </tr>
            </thead>
            <tbody>
                <tr>
                    <td class="border px-4 py-2">Web</td>
                    <td class="border px-4 py-2">{{ totalWeb }}</td>
                </tr>
                <!-- <tr>
                    <td class="border px-4 py-2">Diễn Đàn</td>
                    <td class="border px-4 py-2">{{ totalForum }}</td>
                </tr> -->
                <tr>
                    <td class="border px-4 py-2">Tin tức</td>
                    <td class="border px-4 py-2">{{ totalNew }}</td>
                </tr>
                <tr>
                    <td class="border px-4 py-2">Facebook</td>
                    <td class="border px-4 py-2">{{ totalFace }}</td>
                </tr>
                <tr>
                    <td class="border px-4 py-2">Video</td>
                    <td class="border px-4 py-2">{{ totalVideo }}</td>
                </tr>
            </tbody>
        </table>
    </div>
</template>
  



<script>
import axios from 'axios';

export default {
    name: 'ListSource',

    data() {
        return {
            totalVideo: 0,
            // totalForum: 0,
            totalWeb: 0,
            totalFace: 0,
            totalNew: 0,
        };
    },
    created() {
        this.fetchSources()
    },
    methods: {
        async fetchSources() {
            try {
                const response = await axios.request({
                    methods: 'GET',
                    url: 'http://localhost:6006/api/v1/common/sources',
                });
                console.log({ response })
                this.totalNew = response.data.data.countNew;
                this.totalVideo = response.data.data.countVideo;
                // this.totalForum = response.data.data.countForum;
                this.totalWeb = response.data.data.countWeb;
                this.totalFace = response.data.data.countFace;
            } catch (error) {
                console.error('Error fetching sources:', error);
            }
        },
    }
};
</script>
<style scoped>
.item-container {
    display: flex;
    flex-direction: column;
}

.item {
    margin-bottom: 10px;
    /* hoặc bất kỳ giá trị khoảng cách nào phù hợp */
}
</style>
