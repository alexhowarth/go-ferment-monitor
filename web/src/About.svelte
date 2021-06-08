<script>
  import { onMount } from "svelte";

  import Dialog, { Title, Content, Actions } from "@smui/dialog";
  import Button, { Label } from "@smui/button";

  export function openAbout() {
    dialog.open();
  }

  let dialog;
</script>

<Dialog
  bind:this={dialog}
  aria-labelledby="dialog-title"
  aria-describedby="dialog-content">
  <!-- on:MDCDialog:closed={closeHandler}> -->
  <Title id="dialog-title">About</Title>
  <Content id="dialog-content">(c) Alex Howarth, 2020</Content>
  <Actions>
    <Button>
      <Label>OK</Label>
    </Button>
  </Actions>
</Dialog>
