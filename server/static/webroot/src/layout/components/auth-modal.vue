<script setup lang=ts>
//base
import { ref, onMounted } from 'vue'
import { storeToRefs } from 'pinia';
import { useAuthStore } from '@/stores/auth';
import { NModal, NCard, NInput, NButton } from 'naive-ui';
//icons

//store
const AuthStore = useAuthStore()
const { showAuthModal } = storeToRefs(AuthStore)
//ref
const key = ref<string>("")
//hook

//fn
const submitKey = async () => {
    await AuthStore.auth(key.value)
}
</script>
<template>
    <n-modal v-model:show="showAuthModal">
        <n-card class="auth-cont" title="访问密钥">
            <n-input v-model:value="key" placeholder="访问密钥" class="auth-input"></n-input>
            <n-button @click="submitKey" type="info">确认</n-button>
        </n-card>
    </n-modal>
</template>
<style>
.auth-cont {
    width: 50%;
    position: absolute;
    transform: translateX(-50%);
    left: 50%;
}

.auth-input {
    margin-bottom: 10px;
}

@media (max-width: 1000px) {
    .auth-cont {
        width: 70%;
    }
}

@media (max-width: 500px) {
    .auth-cont {
        width: 80%;
    }
}
</style>