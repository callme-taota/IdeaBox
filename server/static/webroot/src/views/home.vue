<script setup lang="ts">
// base
import { ref, onMounted } from 'vue'
import { useMessage, NIcon, NDrawer, NDrawerContent, NInput, NSelect, NButton, type DrawerPlacement } from 'naive-ui'
import Idea_Card from '@/components/idea_card.vue'
//icon
import { Add as addicon } from '@vicons/ionicons5';
//api
import { GetIdea, DeleteIdea, CreateIdea, UpdateIdea } from '@/api/idea';
import { GetUserList } from '@/api/user';
//utils
import type { FullIdeaItem, UserItem, Idea, IdeaWithID } from '@/utils/interface';
//store
const Message = useMessage()
//ref
const ideaList = ref<FullIdeaItem[]>([])
const displayDrawer = ref<boolean>(false)
const placement = ref<DrawerPlacement>("right")
const editDisplayDrawer = ref<boolean>(false)
const editplacement = ref<DrawerPlacement>("right")
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
    if (obj.Name == "" || obj.Author > 2 || obj.Author < 1 || obj.Description == "") {
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
</script>
<template>
    <div class="page-cont">
        <div class="waterfall-container">
            <Idea_Card v-for="idea in ideaList" :color="idea.Color" :user-name="idea.Name" :idea-name="idea.IdeaName"
                :idea-description="idea.Description" :create-at="idea.CreatedAt" @edit="editIdea(idea)" />
        </div>
        <div class="add-btn" @click="showDrawer">
            <n-icon size="28">
                <addicon />
            </n-icon>
        </div>
        <n-drawer v-model:show="displayDrawer" :placement="placement" default-width="500" default-height="350">
            <n-drawer-content title="我有一个点子💡">
                <div>
                    点子
                    <n-input v-model:value="ideaItem.Name" placeholder="点子" maxlength="30" show-count
                        clearable></n-input>
                </div>
                <div>
                    谁
                    <n-select v-model:value="ideaItem.Author" placeholder="谁呀" :options="userList" label-field="Name"
                        value-field="ID" filterable></n-select>
                </div>
                <div>
                    细说细说
                    <n-input v-model:value="ideaItem.Description" placeholder="细说细说" type="textarea"
                        clearable></n-input>
                </div>
                <n-button @click="submitIdea">想完了</n-button>
            </n-drawer-content>
        </n-drawer>
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