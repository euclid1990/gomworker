<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class Workers extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('workers', function (Blueprint $table) {
            $table->increments('id');
            $table->enum('status', ['running', 'stopping'])->default('stopping');
            $table->string('queue')->default('default');
            $table->boolean('once')->default(false);
            $table->integer('delay')->default(0);
            $table->boolean('force')->default(false);
            $table->integer('memory')->default(128);
            $table->integer('sleep')->default(3);
            $table->integer('timeout')->default(60);
            $table->integer('tries')->default(0);
            $table->timestamp('started_at')->nullable();
            $table->timestamps();
            $table->softDeletes();
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('workers');
    }
}
