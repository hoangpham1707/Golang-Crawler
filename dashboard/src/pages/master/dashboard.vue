<template>
    <div class="w-screen h-screen flex">
        <!-- Side bar -->
        <div class="w-[400px] h-full bg-white-200">
            <div class="h-[50px] bg-white-900 flex justify-start items-center">
                <div class="px-[20px]">
                    <h3 class="font-bold text-xl">
                        <router-link to="/">Admin DashBoard Crawler</router-link>
                    </h3>
                </div>
            </div>
            <div class="h-[calc(100vh-60px)] bg-rgba(0, 0, 0, 0) py-[20px]">
                <div class="px-[20px] flex flex-col justify-between space-y-[10px]">
                    <div class="h-auto">
                        <div class="h-[50px] flex items-center">
                            <router-link to="/"><i class="fas fa-home"></i><span @mouseover="highlightHome"
                                    @mouseleave="unhighlightHome"> Home</span> </router-link>
                        </div>
                        <div>
                            <div class="h-[50px] flex items-center" @click="toggleSource()">
                                <i class="fas fa-link"></i>
                                <router-link to="/source"><span @mouseover="highlightHome" @mouseleave="unhighlightHome">
                                        Source</span> </router-link>

                            </div>
                            <div v-show="showSource">
                                <div class="h-[50px] flex items-center">
                                    <select id="selectedLabelSource" v-model="selectedLabelSource"
                                        @change="navigateToLabelSource">
                                        <option value=""></option>
                                        <option value="newSource">Tin tức({{ totalNew }})</option>
                                        <option value="webSource">Web({{ totalWeb }})</option>
                                        <!-- <option value="forumSource">Diễn đàn</option> -->
                                        <option value="faceSource">Facebook({{ totalFace }})</option>
                                        <option value="videoSource">Video({{ totalVideo }})</option>
                                    </select>
                                </div>
                            </div>

                            <div>
                                <div class="h-[50px] flex items-center">
                                    <i class="fas fa-link"></i>
                                    <router-link to="/urlNext"><span @mouseover="highlightHome"
                                            @mouseleave="unhighlightHome">
                                            Link Next</span> </router-link>
                                </div>
                            </div>

                            <div class="h-[50px] flex items-center">
                                <router-link to="/category"><i class="fas fa-folder"></i><span @mouseover="highlightHome"
                                        @mouseleave="unhighlightHome"> Category</span> </router-link>
                            </div>

                            <div class="h-[50px] flex items-center">
                                <router-link to="/keyword"><i class="fas fa-key"></i> <span @mouseover="highlightHome"
                                        @mouseleave="unhighlightHome"> Keyword</span> </router-link>
                            </div>
                            <div class="h-[50px] flex items-center">
                                <router-link to="/label"><i class="fas fa-tag"></i> <span @mouseover="highlightHome"
                                        @mouseleave="unhighlightHome">Label</span> </router-link>
                            </div>

                            <div>
                                <div class="h-[50px] flex items-center">
                                    <router-link to="/addLabels" @click="toggleContent('none')">
                                        <i class="fas fa-tag"></i> <span @mouseover="highlightHome"
                                            @mouseleave="unhighlightHome">Add Label</span>
                                    </router-link>
                                </div>

                                <div v-show="showContent">
                                    <div class="h-[50px] flex items-center">
                                        <select id="selectLabel" v-model="selectedLabel" @change="navigateToLabel">
                                            <option value="addLabels"></option>
                                            <option value="faceToiYeuPtit">Tôi Yêu PTIT</option>
                                            <option value="faceCfs247">PTIT Confessions 24/7</option>
                                            <option value="faceNganHangPtit">Ngân hàng PTIT</option>
                                            <option value="faceHmOfPtit">Humans of PTIT</option>
                                            <option value="faceKtxCfs">Kí Túc Xá PTIT Confession</option>
                                            <option value="faceCotSongPtit">Cột sống Ptit</option>
                                            <option value="faceCfs">PTIT Confessions</option>
                                            <option value="faceAnti">Hội anti HV công nghệ BƯU ĐIỆN</option>
                                            <option value="faceDramaPtit">Drama PTIT</option>
                                            <option value="facePtitNews">PTIT News</option>
                                            <option value="faceD21">D21 PTIT Hà Nội</option>
                                            <option value="faceClbTachmon">CLB Tạch môn PTIT</option>

                                            <option value="web">Web</option>
                                        </select>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <!-- Side bar -->

        <!-- Main -->
        <div class="w-[full] h-full bg-gray-400">
            <!-- <div class="h-[50px] bg-gray-100 flex items-center shadow-sm px-[20px] w-full py-[10px] z-10 border-n">Header
            </div> -->
            <div class="h-[50px] bg-gray-100 flex justify-center items-center">

                <!-- <input type="text"
                    class="w-[66.67%] p-2 rounded-md bg-white border-2 border-gray-300 focus:outline-none focus:border-blue-500"
                    placeholder="Tìm kiếm..." /> -->

            </div>

            <div class="h-[calc(100vh-50px)] bg-gray-50 p-[20px]">
                <div class="border border-gray-300 rounded-md p-[20px] h-full">
                    <router-view></router-view>
                </div>
            </div>
        </div>
        <!-- Main -->

    </div>
</template>
  

<script>
import axios from 'axios';
export default {
    name: 'DashBoard',
    data() {
        return {
            totalVideo: 0,
            // totalForum: 0,
            totalWeb: 0,
            totalFace: 0,
            totalNew: 0,
            showContent: false,
            showSource: false,
            selectedLabel: '',
            selectedLabelSource: '',
            labels: [
                { value: 'addLabels', account_id: '' },
                { value: 'faceToiYeuPtit', account_id: '656032122d8638c72626a67e' },
                { value: 'faceCfs247', account_id: '656031bf71ab8c916cc177a4' },
                { value: 'faceNganHangPtit', account_id: '6560317eb1f1ff0adc9fa847' },
                { value: 'faceHmOfPtit', account_id: '65603142f5f12d7287bf6e86' },
                { value: 'faceKtxCfs', account_id: '656030f82a5d479c36d3f088' },
                { value: 'faceCotSongPtit', account_id: '6560307a0f8c9d3ba12e8375' },
                { value: 'faceCfs', account_id: '65603005fbefafb306e0cc8b' },
                { value: 'faceAnti', account_id: '65602f9b75d60a8b5d6be60b' },
                { value: 'faceDramaPtit', account_id: '65602e7c646c9b4cc4bff677' },
                { value: 'facePtitNews', account_id: '65647e1495b6a5f26f5921cb' },
                { value: 'faceD21', account_id: '6565cd469ecf2b9c2ea7255e' },
                { value: 'faceClbTachmon', account_id: '6565c74b9ecf2b9c2ea72467' },
                { value: 'web', account_id: '' }
            ],
            labelSource: [
                { value: 'newSource', typrSource: 'New' },
                { value: 'webSource', typrSource: 'Web' },
                // { value: 'forumSource', typrSource: 'Forum' },
                { value: 'faceSource', typrSource: 'Facebook' },
                { value: 'videoSource', typrSource: 'Video' },
            ],
        };
    },
    created() {
        //this.fetchMentions()
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
        navigateToLabel() {
            const selectedOption = this.labels.find(option => option.value === this.selectedLabel);
            const account_id = selectedOption ? selectedOption.account_id : '';

            if (this.selectedLabel == "web") {
                this.$router.push({
                    path: `/web`,
                    query: { account_id },
                });
            } else {
                this.$router.push({
                    path: `/face`,
                    query: { account_id },
                });
            }

        },
        navigateToLabelSource() {
            const selectedOption = this.labelSource.find(option => option.value === this.selectedLabelSource);
            const typeSource = selectedOption ? selectedOption.typrSource : '';

            this.$router.push({
                path: `/typeSource`,
                query: { typeSource }
            })

        },
        highlightHome(event) {
            event.target.style.backgroundColor = 'green'; // Đổi màu nền khi lăn chuột qua
        },
        unhighlightHome(event) {
            event.target.style.backgroundColor = ''; // Đổi lại màu nền khi lăn chuột ra
        },
        toggleContent() {
            this.showContent = !this.showContent;
            if (this.showContent) {
                event.target.style.backgroundColor = 'green';
            } else {
                event.target.style.backgroundColor = '';
            }
        },
        toggleSource() {
            this.showSource = !this.showSource;
            if (this.showSource) {
                event.target.style.backgroundColor = 'green';
            } else {
                event.target.style.backgroundColor = '';
            }
        },

    }
}

</script>

<style>
span {
    display: inline-block;
    transition: background-color 0.5s;
    padding: 5px;
}

/* Kiểu mặc định khi không lăn chuột */
span:hover {
    background-color: rgb(70, 70, 60);
}
</style>
