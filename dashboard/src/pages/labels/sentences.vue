<!-- ListSource.vue -->
<template>
    <div class="container mx-auto p-4">
        <button @click="exit()" class="ml-2 bg-blue-500 text-white px-2 py-1 rounded">Trở lại</button>
        <h1 class="text-2xl font-semibold mb-4">List Sentences</h1>
        <ul>
            <li v-for="entry in paginatedEntries" :key="entry._id" class="mb-2 flex items-center justify-between">
                <div class="entry-container">
                    <p>{{ entry.LabelId }}</p>
                    <p>{{ entry.Name }}</p>
                </div>
                <div class="flex items-center">
                    <select v-model="entry.selectedLabel">
                        <option v-for="label in labels" :key="label.slug" :value="label.slug">{{ label.name }}</option>
                    </select>
                    <button @click="assignLabel(entry._id, entry.selectedLabel, entry.sentenceIndex)"
                        class="ml-2 bg-green-500 text-white px-2 py-1 rounded">Gán Nhãn</button>
                </div>
            </li>
        </ul>
        <!-- Nút phân trang -->
        <div class="pagination mt-4 flex items-center justify-center">
            <button @click="changePage(-1)" :disabled="currentPage === 1">Previous</button>
            <span class="mx-2">Page {{ currentPage }} of {{ totalPages }}</span>
            <button @click="changePage(1)" :disabled="currentPage === totalPages">Next</button>
        </div>
    </div>
</template>
  
  
<script>
import axios from 'axios';


export default {
    name: 'ListSentences',
    data() {
        return {
            entries: [],
            currentPage: 1, // Trang hiện tại
            pageSize: 10,
            labels: [], // Thêm mảng này để lưu trữ các nhãn
            selectedLabel: '', // Thêm biến này để lưu trữ nhãn đã chọn
        };
    },
    computed: {
        paginatedEntries() {
            const startIndex = (this.currentPage - 1) * this.pageSize;
            const endIndex = startIndex + this.pageSize;
            return this.entries.slice(startIndex, endIndex);
        },
        totalPages() {
            return Math.ceil(this.entries.length / this.pageSize);
        },
    },
    created() {
        this.fetchSources();
        this.fetchLabels();
    },
    methods: {
        async fetchSources() {
            try {
                const response = await axios.request({
                    method: 'GET',
                    url: `http://localhost:6006/api/v1/common/mentions/${this.$route.params.entryId}/sentences`,
                });
                console.log({ response });
                this.entries = response.data.data.entry;

                this.entries.forEach((entry, index) => {
                    entry.sentenceIndex = index + (this.currentPage - 1) * this.pageSize;
                });
            } catch (error) {
                console.error('Error fetching sources:', error);
            }
        }, async assignLabel(entryId, selectedLabel, sentenceIndex) {
            try {
                const response = await axios.put(`http://localhost:6006/api/v1/common/mentions/${this.$route.params.entryId}/assign-label-sentences?label_name=${selectedLabel}&sentence_index=${sentenceIndex}`);
                const mention = response.data.data.mention;

                const label = this.labels.find(label => label.slug === mention);
                console.log("label: ", label)
                if (label) {
                    const entry = this.entries.find(entry => entry._id === entryId);
                    entry.label = label;
                    console.log("nameMention: ", entry.label.name)
                }
                await this.fetchSources();
            } catch (error) {
                console.error('Lỗi khi gán nhãn:', error);
            }
        },
        async fetchLabels() {
            try {
                const response = await axios.get('http://localhost:6006/api/v1/common/labels');
                this.labels = response.data.data.entries;
                this.entries.forEach(entry => {
                    entry.selectedLabel = this.labels[0]?.name || ''; // Default to the first label
                });
            } catch (error) {
                console.error('Lỗi khi lấy danh sách nhãn:', error);
            }
        },
        async exit() {
            const account_id = this.$route.query.account_id;
            if (account_id != '') {
                this.$router.push({ path: `/face`, query: { account_id }, },);
            } else {
                this.$router.push({ path: `/web` },);
            }
        },
        async onPageChange(page) {
            this.currentPage = page;
            await this.fetchSources();
        },
        changePage(offset) {
            const newPage = this.currentPage + offset;
            if (newPage > 0 && newPage <= this.totalPages) {
                this.currentPage = newPage;
            }
        },
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
</style>
