<script>
  import { onMount } from "svelte";

  import TabBar from "@smui/tab-bar";
  import Tab, { Icon as TabIcon, Label as TabLabel } from "@smui/tab";
  import Switch from "@smui/switch";
  import FormField from "@smui/form-field";
  import Textfield from "@smui/textfield";
  import HelperText from "@smui/textfield/helper-text/index";
  import Paper, {
    Title as PaperTitle,
    Subtitle as PaperSubtitle,
    Content as PaperContent
  } from "@smui/paper";

  import Button, {
    Group as ButtonGroup,
    GroupItem as ButtonGroupItem,
    Label as ButtonLabel,
    Icon as ButtonIcon
  } from "@smui/button";

  import { createForm } from "svelte-forms-lib";
  import * as yup from "yup";

  let elevation = 3;
  let active;
  //let selected = false;

  const mqttEndpoint = "http://localhost:8080/v1/mqtt/";

  const { form, errors, state, handleChange, handleSubmit } = createForm({
    initialValues: {
      id: 1,
      url: "",
      enabled: false,
      username: "",
      password: ""
    },
    validationSchema: yup.object().shape({
      enabled: yup.boolean().required(),
      url: yup
        .string()
        .matches(
          /^((mqtts?|tcp):)?\/\/(((([a-z]|\d|-|\.|_|~|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])|(%[\da-f]{2})|[!\$&'\(\)\*\+,;=]|:)*@)?(((\d|[1-9]\d|1\d\d|2[0-4]\d|25[0-5])\.(\d|[1-9]\d|1\d\d|2[0-4]\d|25[0-5])\.(\d|[1-9]\d|1\d\d|2[0-4]\d|25[0-5])\.(\d|[1-9]\d|1\d\d|2[0-4]\d|25[0-5]))|((([a-z]|\d|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])|(([a-z]|\d|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])([a-z]|\d|-|\.|_|~|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])*([a-z]|\d|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])))\.)+(([a-z]|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])|(([a-z]|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])([a-z]|\d|-|\.|_|~|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])*([a-z]|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])))\.?)(:\d*)?)(\/((([a-z]|\d|-|\.|_|~|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])|(%[\da-f]{2})|[!\$&'\(\)\*\+,;=]|:|@)+(\/(([a-z]|\d|-|\.|_|~|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])|(%[\da-f]{2})|[!\$&'\(\)\*\+,;=]|:|@)*)*)?)?(\?((([a-z]|\d|-|\.|_|~|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])|(%[\da-f]{2})|[!\$&'\(\)\*\+,;=]|:|@)|[\uE000-\uF8FF]|\/|\?)*)?(\#((([a-z]|\d|-|\.|_|~|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])|(%[\da-f]{2})|[!\$&'\(\)\*\+,;=]|:|@)|\/|\?)*)?$/i,
          "Not a valid URL"
        )
        .required(),
      username: yup.string().required(),
      password: yup.string().required()
    }),
    onSubmit: values => {
      doPut(JSON.stringify(values));
    }
  });

  onMount(async () => {
    const res = await fetch(mqttEndpoint);
    if (await res.ok) {
      $form = await res.json();
    }
  });

  async function doPut(body) {
    const res = await fetch(mqttEndpoint, {
      method: "PUT",
      body: body
    });
  }

  //   let urlField;
  //   let urlLabel;

  //   $: if ($errors.url) {
  //     urlLabel = $errors.url;
  //     urlField.focus();
  //   } else {
  //     urlLabel = "URL";
  //   }
</script>

<!-- <style>
  .paper-container {
    background-color: #f8f8f8;
    padding: 36px 18px;
  }
  * :global(.paper-demo) {
    margin: 0 auto;
    max-width: 600px;
  }
</style> -->
<!-- <section> -->
<div>
  <TabBar tabs={['MQTT', 'BrewFather']} let:tab bind:active>
    <Tab {tab}>
      <TabLabel>{tab}</TabLabel>
    </Tab>
  </TabBar>
</div>
<!-- </section> -->

<div style="padding-top: 50px;">
  {#if active == 'MQTT'}
    <!-- <Paper {elevation}>
      <PaperContent> -->
    <div>
      <FormField>
        <Switch bind:checked={$form.enabled} id="enabled" name="enabled" />
        <span slot="label">Enabled</span>
      </FormField>
    </div>
    <br />
    <div>
      <Textfield
        dense
        variant="outlined"
        bind:value={$form.url}
        label="URL"
        id="url"
        name="url"
        invalid={$errors.url} />
      <HelperText id="helper-text-dense-a">URL of the MQTT server</HelperText>
      <br />
      <Textfield
        dense
        variant="outlined"
        bind:value={$form.username}
        label="Username"
        id="username"
        name="username"
        invalid={$errors.username} />
      <HelperText id="helper-text-dense-a">MQTT user name</HelperText>
      <br />
      <Textfield
        dense
        variant="outlined"
        bind:value={$form.password}
        label="Password"
        id="password"
        name="password"
        invalid={$errors.password} />
      <HelperText id="helper-text-dense-a">MQTT password</HelperText>
      <br />
      <Button on:click={handleSubmit} variant="raised" color="secondary">
        <ButtonLabel>Save</ButtonLabel>
      </Button>
    </div>
    <!-- </PaperContent>
    </Paper> -->
  {/if}
  {#if active == 'BrewFather'}
    Brewfather here
    <!-- <Paper {elevation}>
      <PaperContent>BrewFather config here!</PaperContent>
    </Paper> -->
  {/if}
</div>

<!-- </div> -->
<!-- 
<div
  style="display: flex; flex-wrap: wrap; justify-content: center; align-items:
  center; padding-top: 100px;">
  <Paper elevation="100" style="width: 400px;">
    <PaperContent>
      <div>
        <FormField>
          <Switch bind:checked={selected} />
          <span slot="label">Enabled</span>
        </FormField>
      </div>
      <br />
      <div>
        <Textfield dense variant="outlined" bind:value={url} label="URL" />
        <HelperText id="helper-text-dense-a">
          URL for your MQTT server
        </HelperText>
        <br />
        <Textfield
          dense
          variant="outlined"
          bind:value={username}
          label="Username" />
        <HelperText id="helper-text-dense-a">MQTT user name</HelperText>
        <br />
        <Textfield
          dense
          variant="outlined"
          bind:value={password}
          label="Password" />
        <HelperText id="helper-text-dense-a">MQTT password</HelperText>
        <br />
        <Button variant="raised">
          <ButtonLabel>Save</ButtonLabel>
        </Button>
      </div>
    </PaperContent>
  </Paper>
</div> -->
