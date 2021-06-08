<script>
  import { onMount } from "svelte";
  import DataTable, { Head, Body, Row, Cell } from "@smui/data-table";
  import Textfield from "@smui/textfield";
  import Button, {
    Group as ButtonGroup,
    GroupItem as ButtonGroupItem,
    Label as ButtonLabel,
    Icon as ButtonIcon
  } from "@smui/button";

  import { createForm } from "svelte-forms-lib";
  import * as yup from "yup";
  import Checkbox from "@smui/checkbox";
  import HelperText from "@smui/textfield/helper-text/index";

// THIS	import Device from './Device.svelte';


  const tiltEndpoint = "http://localhost:8080/v1/tilt/";

  let tilts = [] // use this as can't use #$form when building table (forms are now in a separate file)

  onMount(async () => {
    const res = await fetch(tiltEndpoint);
    //$form.tilts = await res.json();
    tilts = await res.json();
  });


  async function doPut(body) {
    const res = await fetch(tiltEndpoint, {
      method: "PUT",
      body: body
    });
  }

  // move this into a separate 'module' so that each device is a form
  const { form, errors, state, handleChange, handleSubmit } = createForm({
    initialValues: {
      /*
      Colour    tilt.Colour `gorm:"unique_index;auto_increment:false;not null" json:"colour"` // not sure about auto_increment here
	    Enabled   bool        `json:"enabled"`
	    Name      string      `json:"name"`
	    OG        float64     `json:"og"`
	    Consumers Consumers   `gorm:"foreignkey:id" json:"consumers"` //`gorm:"no null"`*/
      //tilts: [] // this is wrong? i want individual forms? or?
      id: "",
      name: "",
      enabled: false,
      og: 0,
      colour: 0,
      consumers: {
        mqtt: false,
        brewfather: false,
      }
    },
    validationSchema: yup.object().shape({
      tilts: yup.array().of(
        yup.object().shape({
          name: yup.string(),
          og: yup.number()
        })
      )
    }),
    onSubmit: values => {
      doPut(JSON.stringify(values));
    }
  });

  function tiltCol(val) {
    return [
      "Red",
      "Green",
      "Black",
      "Purple",
      "Orange",
      "Blue",
      "Yellow",
      "Pink"
    ][val];
  }
</script>

<div style="padding-top: 50px;">
  <DataTable table$aria-label="Devices">
    <Head>
      <Row>
        <Cell>Type</Cell>
        <Cell>Description</Cell>
        <Cell>Name</Cell>
        <Cell>OG</Cell>
        <Cell>MQTT</Cell>
        <Cell>BrewFather</Cell>
        <Cell />
        <Cell />
      </Row>
    </Head>
    <Body>
      <!-- {#each $form.tilts as tilt, j} -->
      <!-- put each of these in a deviceform svelte module thing - move all of the form code into this so that they are all separate forms - job done! -->
      {#each tilts as tilt, j}
      <!-- this becomes <Device /> with arguments passed in -->
        <Row>
          <Cell>Tilt</Cell>
          <Cell>{tiltCol(tilt.colour)}</Cell>
          <Cell>
            <!-- <input bind:value={$form.tilts[j].name} name="name" /> -->
            <input bind:value={$form.name} name="name" />
          </Cell>
          <Cell>
            <input
              style="width: 50px"
              bind:value={$form.og}
              name="og" />
          </Cell>
          <Cell>
            <input
              name="mqtt"
              type="checkbox"
              bind:value={$form.consumers.mqtt}
              bind:checked={$form.consumers.mqtt} />
          </Cell>
          <Cell>
            <input
              name="brewfather"
              type="checkbox"
              bind:checked={$form.consumers.brewfather}
              bind:value={$form.consumers.brewfather} />
          </Cell>
          <Cell>
            <Button on:click={handleSubmit} variant="raised" color="secondary">
              <!-- <Button on:click={() => handleSubmit($form.tilts[j])} variant="raised" color="secondary"> -->
              <ButtonLabel>Save</ButtonLabel>
            </Button>
          </Cell>
          <Cell>
            <Button variant="raised" color="primary">
              <ButtonLabel>Delete</ButtonLabel>
            </Button>
          </Cell>
        </Row>
      {/each}
    </Body>
  </DataTable>
</div>
