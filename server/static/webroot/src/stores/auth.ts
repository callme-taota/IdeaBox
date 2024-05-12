import { Auth } from "@/api/common";
import { defineStore } from "pinia";
import { ref } from "vue";

export const useAuthStore = defineStore("authStore", () => {
    const showAuthModal = ref<boolean>(false)

    const auth = async (key: string) => {
        await Auth({ 'Key': key })
        showAuthModal.value = false
        setTimeout(() =>{
            window.location.reload()
        },1000)
    }

    return { showAuthModal, auth }
})