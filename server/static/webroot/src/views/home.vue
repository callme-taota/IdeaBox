<script setup lang="ts">
// base
import { ref, onMounted, nextTick } from 'vue'
import { useMessage, NIcon, NDropdown, type DrawerPlacement } from 'naive-ui'
import Idea_Card from '@/components/idea_card.vue'
import Idea_add from '@/components/idea_add.vue';
import Idea_edit from '@/components/idea_edit.vue';
import Comment from '@/components/comment.vue';
//icon
import { Add as addicon } from '@vicons/ionicons5';
//api
import { GetIdea } from '@/api/idea';
import { GetUserList } from '@/api/user';
//utils
import type { FullIdeaItem, UserItem, IdeaWithID } from '@/utils/interface';
//store
const Message = useMessage()
//ref
const ideaList = ref<FullIdeaItem[]>([])

const placement = ref<DrawerPlacement>("right")
const addDrawerDisplay = ref<boolean>(false)
const editDrawerDisplay = ref<boolean>(false)
const commentDrawerDisplay = ref<boolean>(false)

const editItem = ref<IdeaWithID>({
    ID: 0,
    Name: '',
    Description: '',
    Author: 0
})

const targetIdeaID = ref<number>(0)
const userList = ref<UserItem[]>([])

const showDown = ref<boolean>(false)
const downX = ref<number>(0)
const downY = ref<number>(0)
const menuOptions = ref([
    {
        label: '新增',
        key: 'add'
    },
])
//hook
onMounted(async () => {
    let res = await GetIdea({})
    if (res.ok) {
        ideaList.value = res.data.Ideas
    } else {
        Message.error("网络出了点问题诶")
    }
    let user_res = await GetUserList({})
    if (user_res.ok) {
        userList.value = user_res.data
    } else {
        Message.error("网络出了点问题诶")
    }
})
//fn
const getIdea = async () => {
    let res = await GetIdea({})
    if (res.ok) {
        ideaList.value = res.data.Ideas
    } else {
        Message.error("网络出了点问题诶")
    }
}

const showAdd = () => {
    let windowWidth = window.innerWidth
    if (windowWidth <= 900) {
        placement.value = "bottom"
    } else {
        placement.value = "right"
    }
    addDrawerDisplay.value = true
}

const showEditIdea = async (idea: FullIdeaItem) => {
    editItem.value = {
        "ID": idea.IdeaID,
        "Name": idea.IdeaName,
        "Description": idea.Description,
        "Author": idea.Author
    }
    let windowWidth = window.innerWidth
    if (windowWidth <= 900) {
        placement.value = "bottom"
    } else {
        placement.value = "right"
    }
    editDrawerDisplay.value = true
}

const showComment = async (ideaID: number) => {
    targetIdeaID.value = ideaID
    let windowWidth = window.innerWidth
    if (windowWidth <= 900) {
        placement.value = "bottom"
    } else {
        placement.value = "right"
    }
    commentDrawerDisplay.value = true
}

const handleContextMenu = async (e: MouseEvent) => {
    e.preventDefault()
    showDown.value = false
    nextTick().then(() => {
        showDown.value = true
        downX.value = e.clientX
        downY.value = e.clientY
    })
}

const onClickoutside = () => {
    showDown.value = false
}

const handleMenuSelect = (key: string | number) => {
    showDown.value = false
    if (key == "add") {
        showAdd()
    }
}

</script>
<template>
    <div class="page-cont" @contextmenu="handleContextMenu">
        <!-- main part -->
        <div class="waterfall-container">
            <Idea_Card v-for="idea in ideaList" :color="idea.Color" :user-name="idea.Name" :idea-name="idea.IdeaName"
                :idea-description="idea.Description" :create-at="idea.CreatedAt" @edit="showEditIdea(idea)"
                @comment="showComment(idea.IdeaID)" />
        </div>
        <!-- add button -->
        <div class="add-btn" @click="showAdd">
            <n-icon size="28">
                <addicon />
            </n-icon>
        </div>
        <!-- add drawer -->
        <Idea_add @done="getIdea" v-model:show="addDrawerDisplay" :placement="placement"></Idea_add>
        <!-- edit drawer -->
        <Idea_edit @done="getIdea" v-model:show="editDrawerDisplay" v-model:editItem="editItem" :placement="placement" :userList="userList"></Idea_edit>
        <!-- comment drawer -->
        <Comment v-model:show="commentDrawerDisplay" :placement="placement" v-model:ideaID="targetIdeaID"></Comment>
        <n-dropdown placement="bottom-start" trigger="manual" :x="downX" :y="downY" :options="menuOptions"
            :show="showDown" :on-clickoutside="onClickoutside" @select="handleMenuSelect"></n-dropdown>
    </div>
</template>
<style>
.waterfall-container {
    column-count: 3;
    column-gap: 10px;
}

.add-btn {
    position: fixed;
    right: 120px;
    bottom: 80px;
    width: 40px;
    height: 40px;
    border-radius: 50%;
    display: flex;
    justify-content: center;
    align-items: center;
    background: var(--add-btn-background);
    backdrop-filter: blur(4px);
    transition: 0.3s;
    cursor: pointer;
}

.add-btn:hover {
    background: var(--second-highlight-color);
    box-shadow: 0 0 12px var(--second-highlight-color);
    color: var(--color-rev);
}

.comment-card {
    margin: 10px 6px;
    border-radius: 10px;
    padding-bottom: 10px;
    border-bottom: var(--content-border);
}

.comment-card-content {
    margin-top: 6px;
    padding: 2px;
}

.comment-card-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
}

.comment-card-author {
    padding: 1px 6px;
    font-weight: 300;
    border-radius: 6px;
    margin-right: 6px;
}

.comment-card-del {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 24px;
    height: 24px;
    border-radius: 50%;
    background: var(--add-btn-background);
    cursor: pointer;
    transition: 0.3s;
    margin: 0 0 0 8px;
}

.n-drawer-footer {
    display: block !important;
}

@media (max-width:900px) {
    .waterfall-container {
        column-count: 2;
    }

    .add-btn {
        left: 50%;
        transform: translateX(-50%);
        width: 60px;
        height: 60px;
    }
}

@media (max-width:500px) {
    .waterfall-container {
        column-count: 1;
    }
}
</style>