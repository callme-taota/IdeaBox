<script setup lang="ts">
// base
import { ref, onMounted, nextTick } from 'vue'
import { useMessage, NIcon, NDrawer, NDrawerContent, NInput, NSelect, NButton, NDropdown, useDialog, type DrawerPlacement } from 'naive-ui'
import Idea_Card from '@/components/idea_card.vue'
//icon
import { Add as addicon } from '@vicons/ionicons5';
import { Delete as delicon } from '@vicons/carbon'
//api
import { GetIdea, DeleteIdea, CreateIdea, UpdateIdea } from '@/api/idea';
import { CreateComment, DeleteComment, GetComment } from '@/api/comment';
import { GetUserList } from '@/api/user';
//utils
import type { FullIdeaItem, UserItem, Idea, IdeaWithID, CommentItem } from '@/utils/interface';
import { dateToDescription } from '@/utils/time';
//store
const Message = useMessage()
const Dialog = useDialog()
//ref
const ideaList = ref<FullIdeaItem[]>([])
const displayDrawer = ref<boolean>(false)
const placement = ref<DrawerPlacement>("right")
const editDisplayDrawer = ref<boolean>(false)
const editplacement = ref<DrawerPlacement>("right")
const commentDisplayDrawer = ref<boolean>(false)
const commentplacement = ref<DrawerPlacement>("right")
const commentInput = ref<string>("")
const commentWho = ref<number>(1)
const commentList = ref<CommentItem[]>([])
const commentCount = ref<number>(0)
const commentIdeaID = ref<number>(0)
const targetCommentID = ref<number>(0)
const userList = ref<UserItem[]>([])
const ideaItem = ref<Idea>({
    Name: '',
    Description: '',
    Author: 1
})
const editItem = ref<IdeaWithID>({
    ID: 0,
    Name: '',
    Description: '',
    Author: 0
})
const showDown = ref<boolean>(false)
const downX = ref<number>(0)
const downY = ref<number>(0)
const dialogOptions = {
    title: '删除',
    content: '是否删除这条评论?',
    positiveText: '确定',
    negativeText: '不确定',
    onPositiveClick: async () => {
        await deleteComment()
    },
    onNegativeClick: () => {
    }
}
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
const showDrawer = () => {
    let windowWidth = window.innerWidth
    if (windowWidth <= 900) {
        placement.value = "bottom"
    } else {
        placement.value = "right"
    }
    displayDrawer.value = true
}

const submitIdea = async () => {
    let obj = ideaItem.value
    if (obj.Name == "" || obj.Description == "") {
        Message.warning("看看是不是少写了些啥")
    }
    await CreateIdea(obj)
    displayDrawer.value = false
    let res = await GetIdea({})
    if (res.ok) {
        ideaList.value = res.data.Ideas
    } else {
        Message.error("网络出了点问题诶")
    }
    ideaItem.value = {
        Name: "",
        Description: "",
        Author: ideaItem.value.Author,
    }
}

const updateIdea = async () => {
    let obj = editItem.value
    if (obj.Name == "" || obj.Author > 2 || obj.Author < 1 || obj.Description == "") {
        Message.warning("看看是不是少写了些啥")
    }
    await UpdateIdea(obj)
    editDisplayDrawer.value = false
    let res = await GetIdea({})
    if (res.ok) {
        ideaList.value = res.data.Ideas
    } else {
        Message.error("网络出了点问题诶")
    }
    editItem.value = {
        ID: 0,
        Name: "",
        Description: "",
        Author: editItem.value.Author,
    }
}

const editIdea = async (idea: FullIdeaItem) => {
    let obj = {
        "ID": idea.IdeaID,
        "Name": idea.IdeaName,
        "Description": idea.Description,
        "Author": idea.Author
    }
    editItem.value = obj
    let windowWidth = window.innerWidth
    if (windowWidth <= 900) {
        editplacement.value = "bottom"
    } else {
        editplacement.value = "right"
    }
    editDisplayDrawer.value = true
}

const delIdea = async () => {
    let id = editItem.value.ID
    await DeleteIdea({
        "ID": id,
    })
    editDisplayDrawer.value = false
    let res = await GetIdea({})
    if (res.ok) {
        ideaList.value = res.data.Ideas
    } else {
        Message.error("网络出了点问题诶")
    }
    editItem.value = {
        ID: 0,
        Name: "",
        Description: "",
        Author: editItem.value.Author,
    }
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
        showDrawer()
    }
}

const showComment = async (ideaID: number) => {
    commentIdeaID.value = ideaID
    await getComment(ideaID)
    let windowWidth = window.innerWidth
    if (windowWidth <= 900) {
        commentplacement.value = "bottom"
    } else {
        commentplacement.value = "right"
    }
    commentDisplayDrawer.value = true
}

const getComment = async (ideaID: number) => {
    let res = await GetComment({ "IdeaID": ideaID })
    if (res.ok) {
        commentCount.value = res.data.CommentCount
        commentList.value = res.data.CommentList
    }
}

const doComment = async () => {
    let ideaID = commentIdeaID.value
    let content = commentInput.value
    let userID = commentWho.value
    await CreateComment({
        "IdeaID": ideaID,
        "Content": content,
        "UserID": userID
    })
    await getComment(ideaID)
}

const handleDelete = (commentID: number) => {
    targetCommentID.value = commentID
    Dialog.warning(dialogOptions)
}

const deleteComment = async () => {
    let ideaID = commentIdeaID.value
    await DeleteComment({ "CommentID": targetCommentID.value })
    await getComment(ideaID)
}

</script>
<template>
    <div class="page-cont" @contextmenu="handleContextMenu">
        <!-- main part -->
        <div class="waterfall-container">
            <Idea_Card v-for="idea in ideaList" :color="idea.Color" :user-name="idea.Name" :idea-name="idea.IdeaName"
                :idea-description="idea.Description" :create-at="idea.CreatedAt" @edit="editIdea(idea)"
                @comment="showComment(idea.IdeaID)" />
        </div>
        <!-- add button -->
        <div class="add-btn" @click="showDrawer">
            <n-icon size="28">
                <addicon />
            </n-icon>
        </div>
        <!-- add drawer -->
        <n-drawer v-model:show="displayDrawer" :placement="placement" default-width="500" default-height="350">
            <n-drawer-content title="我有一个点子💡">
                <div>
                    点子
                    <n-input v-model:value="ideaItem.Name" placeholder="点子" maxlength="30" show-count
                        clearable></n-input>
                </div>
                <!-- <div>
                    谁
                    <n-select v-model:value="ideaItem.Author" placeholder="谁呀" :options="userList" label-field="Name"
                        value-field="ID" filterable></n-select>
                </div> -->
                <div>
                    细说细说
                    <n-input v-model:value="ideaItem.Description" placeholder="细说细说" type="textarea"
                        clearable></n-input>
                </div>
                <n-button @click="submitIdea">想完了</n-button>
            </n-drawer-content>
        </n-drawer>
        <!-- edit drawer -->
        <n-drawer v-model:show="editDisplayDrawer" :placement="editplacement" default-width="500" default-height="370">
            <n-drawer-content title="点子得改改💡">
                <div>
                    点子
                    <n-input v-model:value="editItem.Name" placeholder="点子" maxlength="30" show-count
                        clearable></n-input>
                </div>
                <div>
                    谁
                    <n-select v-model:value="editItem.Author" placeholder="谁呀" :options="userList" label-field="Name"
                        value-field="ID" filterable></n-select>
                </div>
                <div>
                    细说细说
                    <n-input v-model:value="editItem.Description" placeholder="细说细说" type="textarea"
                        clearable></n-input>
                </div>
                <br>
                <div style="display: flex; justify-content: space-between;">
                    <n-button @click="updateIdea">想完了</n-button>
                    <n-button @click="delIdea" type="error">删了它</n-button>
                </div>
            </n-drawer-content>
        </n-drawer>
        <!-- comment drawer -->
        <n-drawer v-model:show="commentDisplayDrawer" :placement="commentplacement" default-width="500"
            default-height="600" resizable>
            <n-drawer-content :title="'评论💡 ( ' + commentCount + ' ) '">
                <div v-for="comment in commentList" class="comment-card">
                    <div class="comment-card-header">
                        <div class="comment-card-author" :style="{ background: comment.UserColor, color: '#fff' }">
                            {{ comment.UserName }}
                        </div>
                        <div style="display: flex;">
                            <div>
                                {{ dateToDescription(comment.CreatedAt) }}
                            </div>
                            <div class="comment-card-del" @click="handleDelete(comment.ID)">
                                <n-icon>
                                    <delicon />
                                </n-icon>
                            </div>
                        </div>
                    </div>
                    <div class="comment-card-content">
                        {{ comment.Content }}
                    </div>
                </div>
                <template #footer>
                    <div>
                        <n-input v-model:value="commentInput" placeholder="细说细说" type="textarea" clearable
                            maxlength="120"></n-input>
                        <div style="display: flex; margin-top: 12px;">
                            <div style="flex:1"></div>
                            <!-- <n-select v-model:value="commentWho" placeholder="谁呀" :options="userList" label-field="Name"
                                value-field="ID" filterable></n-select> -->
                            <div style="width: 10px;"></div>
                            <n-button @click="doComment">评论</n-button>
                        </div>
                    </div>
                </template>
            </n-drawer-content>
        </n-drawer>
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